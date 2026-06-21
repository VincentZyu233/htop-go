#!/bin/bash
# htop-go installer (Gitee mirror) — supports apt / dnf / Termux
# Usage: curl -fsSL https://gitee.com/vincent-zyu/htop-go/raw/main/scripts/install_gitee.sh | bash
# Install specific version: HTOP_GO_VERSION=v0.1.0 bash -c "$(curl -fsSL https://gitee.com/vincent-zyu/htop-go/raw/main/scripts/install_gitee.sh)"
set -e

OWNER="vincent-zyu"
REPO="htop-go"
API_URL="https://gitee.com/api/v5/repos/${OWNER}/${REPO}/releases/latest"
BIN_NAME="htop-go"

need_cmd() {
  if ! command -v "$1" >/dev/null 2>&1; then
    echo "❌ 缺少必要命令: $1"
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
    echo "❌ 不支持的架构: $ARCH"
    echo "   支持的架构: amd64、arm64、arm"
    echo "   发布页: https://gitee.com/${OWNER}/${REPO}/releases"
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
  echo "❌ 不支持的平台。"
  echo "   本安装脚本支持 apt（Debian/Ubuntu）、dnf（Fedora/RHEL）和 Termux。"
  echo "   发布页: https://gitee.com/${OWNER}/${REPO}/releases"
  exit 1
fi

need_cmd curl
need_cmd install

if [ -n "${HTOP_GO_VERSION:-}" ]; then
  VERSION="$HTOP_GO_VERSION"
  echo "📌 使用指定版本: $VERSION"
else
  echo "📡 正在获取最新版本..."
  VERSION="$(curl -fsSL "$API_URL" | grep '"tag_name"' | head -1 | sed 's/.*"tag_name": *"\([^"]*\)".*/\1/')"
  if [ -z "$VERSION" ]; then
    echo "❌ 从 Gitee API 获取最新版本失败。"
    exit 1
  fi
  echo "📦 最新版本: $VERSION"
fi

ASSET="${BIN_NAME}-${PLATFORM}-${ASSET_ARCH}-${VERSION}"
BASE_URL="https://gitee.com/${OWNER}/${REPO}/releases/download/${VERSION}"
TMP_DIR="$(mktemp -d)"
trap 'rm -rf "$TMP_DIR"' EXIT

echo "🔍 检测到: 架构=${ARCH} 目标架构=${ASSET_ARCH} 包管理器=${PKG_MGR}"
echo "📥 正在下载 ${ASSET}..."
curl -fSL -o "${TMP_DIR}/${BIN_NAME}" "${BASE_URL}/${ASSET}"

if $is_termux; then
  TARGET_DIR="${PREFIX}/bin"
  echo "📦 安装到 ${TARGET_DIR}/${BIN_NAME}..."
  install -Dm755 "${TMP_DIR}/${BIN_NAME}" "${TARGET_DIR}/${BIN_NAME}"
else
  TARGET_DIR="/usr/local/bin"
  echo "📦 安装到 ${TARGET_DIR}/${BIN_NAME}..."
  if command -v sudo >/dev/null 2>&1; then
    sudo install -Dm755 "${TMP_DIR}/${BIN_NAME}" "${TARGET_DIR}/${BIN_NAME}"
  else
    install -Dm755 "${TMP_DIR}/${BIN_NAME}" "${TARGET_DIR}/${BIN_NAME}"
  fi
fi

echo ""
echo "✅ htop-go 安装成功。"
echo "   运行 '${BIN_NAME} --help' 检查安装结果。"
echo ""
echo "   卸载方式:"
if $is_termux; then
  echo "   rm ${PREFIX}/bin/${BIN_NAME}"
else
  echo "   sudo rm -f /usr/local/bin/${BIN_NAME}"
fi
echo ""
echo "   📖 GitHub: https://github.com/VincentZyu233/htop-go"
echo "   📖 Gitee:  https://gitee.com/${OWNER}/${REPO}"
echo ""
echo "   🌐 GitHub 安装脚本:"
echo "   curl -fsSL https://raw.githubusercontent.com/VincentZyu233/htop-go/main/scripts/install.sh | bash"
