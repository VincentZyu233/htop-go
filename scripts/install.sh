#!/bin/bash
# htop-go installer — supports apt / dnf / Termux
# Usage: curl -fsSL https://raw.githubusercontent.com/VincentZyuApps/htop-go/main/scripts/install.sh | bash
# Install specific version: HTOP_GO_VERSION=v0.1.0 bash -c "$(curl -fsSL https://raw.githubusercontent.com/VincentZyuApps/htop-go/main/scripts/install.sh)"
set -e

REPO="VincentZyuApps/htop-go"
API_URL="https://api.github.com/repos/${REPO}/releases/latest"
BIN_NAME="htop-go"

need_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "❌ Missing required command: $1"
    exit 1
  fi
}

is_termux=false
if [ -n "${PREFIX:-}" ] && [ -d "${PREFIX}/bin" ]; then
  is_termux=true
fi

ARCH="$(uname -m)"
case "$ARCH" in
  x86_64|amd64) ASSET_ARCH="amd64" ;;
  aarch64|arm64) ASSET_ARCH="arm64" ;;
  armv7l|armv8l|arm) ASSET_ARCH="arm" ;;
  *)
    echo "❌ Unsupported architecture: $ARCH"
    echo "   Supported: amd64, arm64, arm"
    echo "   Releases: https://github.com/${REPO}/releases"
    exit 1
    ;;
esac

if $is_termux; then
  PLATFORM="android"
  PKG_MGR="termux"
elif command -v apt-get >/dev/null 2>&1; then
  PLATFORM="linux"
  PKG_MGR="apt"
elif command -v dnf >/dev/null 2>&1; then
  PLATFORM="linux"
  PKG_MGR="dnf"
else
  echo "❌ Unsupported platform."
  echo "   This installer supports apt (Debian/Ubuntu), dnf (Fedora/RHEL), and Termux."
  echo "   Releases: https://github.com/${REPO}/releases"
  exit 1
fi

need_cmd curl
need_cmd install

if [ -n "${HTOP_GO_VERSION:-}" ]; then
  VERSION="$HTOP_GO_VERSION"
  echo "📌 Using specified version: $VERSION"
else
  echo "📡 Fetching latest version..."
  VERSION="$(curl -fsSL "$API_URL" | grep '"tag_name"' | head -1 | sed 's/.*"tag_name": *"\([^"]*\)".*/\1/')"
  if [ -z "$VERSION" ]; then
    echo "❌ Failed to fetch latest version from GitHub API."
    exit 1
  fi
  echo "📦 Latest version: $VERSION"
fi

ASSET="${BIN_NAME}-${PLATFORM}-${ASSET_ARCH}-${VERSION}"
BASE_URL="https://github.com/${REPO}/releases/download/${VERSION}"
TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

echo "🔍 Detected: arch=${ARCH} asset_arch=${ASSET_ARCH} pkg_mgr=${PKG_MGR}"
echo "📥 Downloading ${ASSET}..."
curl -fSL -o "${TMP_DIR}/${BIN_NAME}" "${BASE_URL}/${ASSET}"

if $is_termux; then
  TARGET_DIR="${PREFIX}/bin"
  echo "📦 Installing to ${TARGET_DIR}/${BIN_NAME}..."
  install -Dm755 "${TMP_DIR}/${BIN_NAME}" "${TARGET_DIR}/${BIN_NAME}"
else
  TARGET_DIR="/usr/local/bin"
  echo "📦 Installing to ${TARGET_DIR}/${BIN_NAME}..."
  if command -v sudo >/dev/null 2>&1; then
    sudo install -Dm755 "${TMP_DIR}/${BIN_NAME}" "${TARGET_DIR}/${BIN_NAME}"
  else
    install -Dm755 "${TMP_DIR}/${BIN_NAME}" "${TARGET_DIR}/${BIN_NAME}"
  fi
fi

echo ""
echo "✅ htop-go installed successfully."
echo "   Run '${BIN_NAME} --help' to verify."
echo ""
echo "   Uninstall:"
if $is_termux; then
  echo "   rm ${PREFIX}/bin/${BIN_NAME}"
else
  echo "   sudo rm -f /usr/local/bin/${BIN_NAME}"
fi
echo ""
echo "   📖 GitHub: https://github.com/${REPO}"
echo "   📖 Gitee:  https://gitee.com/vincent-zyu/htop-go"
echo ""
echo "   🇨🇳 Gitee mirror:"
echo "   curl -fsSL https://gitee.com/vincent-zyu/htop-go/raw/main/scripts/install_gitee.sh | bash"
