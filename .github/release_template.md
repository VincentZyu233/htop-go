<div align="center">

[![Downloads](https://img.shields.io/github/downloads/__REPO__/__VERSION__/total?style=flat-square&logo=github)](https://github.com/__REPO__/releases/tag/__VERSION__)

</div>

## ✨ Highlights

- 🎛️ Two upper dashboard styles: `classic` and `overview`
- 📋 Optional lower process table via `--table`
- 📱 Cross-platform builds for Windows, Linux, macOS, and Android
- 🤖 Bubble Tea + Lip Gloss + gopsutil stack

## ⬇️ Downloads

| Platform | amd64 / x86_64 | arm64 | arm |
|----------|----------------|-------|-----|
| **Windows** | [Download](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-windows-amd64-__VERSION__.exe) | [Download](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-windows-arm64-__VERSION__.exe) | — |
| **Linux** | [Download](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-linux-amd64-__VERSION__) | [Download](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-linux-arm64-__VERSION__) | — |
| **macOS** | [Download](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-macos-amd64-__VERSION__) | [Download](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-macos-arm64-__VERSION__) | — |
| **Android / Termux** | [Download](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-android-amd64-__VERSION__) | [Download](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-android-arm64-__VERSION__) | [Download](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-android-arm-__VERSION__) |

## 🚀 Quick Start

```bash
# Show help
./htop-go-<target> --help

# Classic upper dashboard only
./htop-go-<target> --style classic

# Overview dashboard with process table
./htop-go-<target> --style overview --table
```

## ⚙️ Common Flags

```text
--style classic|overview
--table
--task
--load
-t, --interval
```

## 📝 Notes

- `--table`, `--task`, and `--load` are opt-in because they cost more sampling work.
- Android builds may differ by architecture because some targets require NDK-based external linking.
- The renderer under `internal/dashboard` stays reusable and separate from the Bubble Tea app shell.
- The single version source is `internal/version/version.go`.
