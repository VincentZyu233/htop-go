# 构建与发布工作流

> **[📖 English](build.md)**
> **[📖 简体中文](build.zh-cn.md)**

## 📋 概述

这套工作流刻意保持简单。

- 版本号只来自 `internal/version/version.go`
- CI 会读取这个文件，并把 `0.1.0` 转成 release tag `v0.1.0`
- commit message 关键词只负责控制“构建”还是“发布”

## 🔑 关键词

| Commit 信息中的关键词 | 构建产物 | GitHub Release |
|----------------------|:---:|:---:|
| `build action` | ✅ | ❌ |
| `build release` | ✅ | ✅ |

说明：

- Pull Request 始终构建
- `workflow_dispatch` 也会构建
- 如果 `main` 上的 commit 没带关键词，则跳过构建与发布

## 🚀 用法示例

```bash
# 仅构建
git commit --allow-empty -m "ci: verify matrix (build action)"

# 构建并创建 GitHub Release
git commit --allow-empty -m "release: ship 0.1.0 (build release)"

# 普通提交，不触发构建 / 发布
git commit -m "docs: update README"
```

## 🏗️ 构建目标

| 平台 | 架构 | 说明 |
|------|:---:|------|
| Windows | `amd64` | `.exe` 二进制 |
| Windows | `arm64` | `.exe` 二进制 |
| Linux | `amd64` | 原生 Go 二进制 |
| Linux | `arm64` | 原生 Go 二进制 |
| macOS | `amd64` | 原生 Go 二进制 |
| macOS | `arm64` | 原生 Go 二进制 |
| Android | `arm64` | 面向 Termux 的目标 |
| Android | `amd64` | 面向模拟器 / x86_64 Android |
| Android | `arm32` | 32 位 ARM Android |

## 📦 流水线阶段

```text
check -> build -> release
```

### `check`

- 解析 commit message
- 读取 `internal/version/version.go`
- 导出 `v<Number>` 作为 workflow 版本号

### `build`

- 使用 `go.mod` 中声明的 Go 版本
- 交叉构建所有配置目标
- 上传产物

### `release`

- 下载构建产物
- 渲染 `.github/release_template.md`
- 使用 `v<Number>` 创建 GitHub Release

## 📌 版本规则

唯一版本来源是：

```text
internal/version/version.go
```

当前形式：

```go
const Number = "0.1.0"
```

当你要发新版本时：

1. 修改 `internal/version/version.go`
2. 提交这个改动
3. 在 commit message 里带上 `build release`

## 🔗 相关文件

- `README.md`
- `README.zh-cn.md`
- `.github/workflows/build.yml`
- `.github/release_template.md`
- `internal/version/version.go`
