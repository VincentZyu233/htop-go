<div align="center">

[![Downloads](https://img.shields.io/github/downloads/__REPO__/__VERSION__/total?style=flat-square&logo=github)](https://github.com/__REPO__/releases/tag/__VERSION__)
[![Windows amd64 | arm64](https://img.shields.io/badge/Windows-amd64_|_arm64-0078D4?style=flat-square&logo=windows&logoColor=white)](https://github.com/__REPO__/releases/tag/__VERSION__)
[![Linux amd64 | arm64](https://img.shields.io/badge/Linux-amd64_|_arm64-FCC624?style=flat-square&logo=linux&logoColor=black)](https://github.com/__REPO__/releases/tag/__VERSION__)
[![macOS amd64 | arm64](https://img.shields.io/badge/macOS-amd64_|_arm64-000000?style=flat-square&logo=apple&logoColor=white)](https://github.com/__REPO__/releases/tag/__VERSION__)
[![Android amd64 | arm64 | arm32](https://img.shields.io/badge/Android-amd64_|_arm64_|_arm32-3DDC84?style=flat-square&logo=android&logoColor=white)](https://github.com/__REPO__/releases/tag/__VERSION__)

</div>

## ✨ Highlights

- 🎛️ Two upper dashboard styles: `classic` and `overview`
- 📋 Optional lower process table via `--table`
- 📱 Cross-platform builds for Windows, Linux, macOS, and Android
- 🤖 Bubble Tea + Lip Gloss + gopsutil stack

## ⬇️ Downloads

| Platform | amd64 / x86_64 | arm64 | arm32 |
|----------|----------------|-------|-------|
| **Windows** | [![windows-amd64](https://img.shields.io/badge/windows-amd64-0078D4.svg?logo=windows&logoColor=white)](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-windows-amd64-__VERSION__.exe) | [![windows-arm64](https://img.shields.io/badge/windows-arm64-0099CC.svg?logo=windows&logoColor=white)](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-windows-arm64-__VERSION__.exe) | — |
| **Linux** | [![linux-amd64](https://img.shields.io/badge/linux-amd64-FCC624.svg?logo=linux&logoColor=black)](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-linux-amd64-__VERSION__) | [![linux-arm64](https://img.shields.io/badge/linux-arm64-E95420.svg?logo=linux&logoColor=white)](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-linux-arm64-__VERSION__) | — |
| **macOS** | [![macos-amd64](https://img.shields.io/badge/macOS-amd64-8E8E93.svg?logo=apple&logoColor=white)](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-macos-amd64-__VERSION__) | [![macos-arm64](https://img.shields.io/badge/macOS-arm64-000000.svg?logo=apple&logoColor=white)](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-macos-arm64-__VERSION__) | — |
| **Android / Termux** | [![android-amd64](https://img.shields.io/badge/android-amd64-8FE388.svg?logo=android&logoColor=white)](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-android-amd64-__VERSION__) | [![android-arm64](https://img.shields.io/badge/android-arm64-168039.svg?logo=android&logoColor=white)](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-android-arm64-__VERSION__) | [![android-arm32](https://img.shields.io/badge/android-arm32-3DDC84.svg?logo=android&logoColor=white)](https://github.com/__REPO__/releases/download/__VERSION__/htop-go-android-arm-__VERSION__) |

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
