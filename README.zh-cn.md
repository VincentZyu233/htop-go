# htop-go

🖥️📊⚡ 一个使用 Go、Bubble Tea、Lip Gloss 和 gopsutil 构建的跨平台 `htop` 风格 TUI。

> **[📖 English](README.md)**
> **[📖 简体中文](README.zh-cn.md)**

[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/VincentZyuApps/htop-go)
[![Gitee](https://img.shields.io/badge/Gitee-C71D23?style=for-the-badge&logo=gitee&logoColor=white)](https://gitee.com/vincent-zyu/htop-go)
[![Go 1.24.6](https://img.shields.io/badge/Go-1.24.6-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![CI](https://img.shields.io/github/actions/workflow/status/VincentZyuApps/htop-go/build.yml?style=for-the-badge&logo=githubactions&logoColor=white&label=CI)](https://github.com/VincentZyuApps/htop-go/actions/workflows/build.yml)

## 🌍 平台目标

- 🪟 Windows
- 🐧 Linux
- 🍎 macOS
- 🤖 Android / Termux

## ✨ 特性

- 🎛️ 提供两种上半区样式：`classic` 和 `overview`
- 📋 通过 `--table` 可选启用下半区进程表
- 🧩 在 `internal/dashboard` 下保留无状态仪表盘渲染层
- 📱 可跨平台构建到 Windows、Linux、macOS 和 Android / Termux
- 🚀 通过 commit message 控制 GitHub Actions 构建与发布流程

## 🚀 本地运行

```powershell
$env:GOCACHE = "$PWD\.gocache"
$env:GOMODCACHE = "$PWD\.gomodcache"
$env:GOPATH = "$PWD\.gopath"
go mod tidy
go run .
```

快速检查：

```powershell
go run . --help
go run . --style classic
go run . --style overview
go run . --table --task
```

## 📥 安装

适用于 Debian / Ubuntu、Fedora / RHEL 和 Termux 的一键安装脚本：

```bash
curl -fsSL https://raw.githubusercontent.com/VincentZyuApps/htop-go/main/scripts/install.sh | bash
htop-go --help
```

中国大陆更友好的 Gitee 镜像：

```bash
curl -fsSL https://gitee.com/vincent-zyu/htop-go/raw/main/scripts/install_gitee.sh | bash
htop-go --help
```

手动指定版本：

```bash
HTOP_GO_VERSION=v0.1.0 bash -c "$(curl -fsSL https://raw.githubusercontent.com/VincentZyuApps/htop-go/main/scripts/install.sh)"
```

> ⚠️ 上面两个 `curl ... | bash` 安装脚本支持 `apt`、`dnf` 和 Termux，本质上会把 release 二进制直接安装到 `PATH` 中。

## ⚙️ CLI

```text
--style classic|overview   选择上半区仪表盘样式
--table                    启用下半区进程表与完整进程采样
--task                     启用任务指标显示
--load                     启用负载指标显示
-t, --interval             刷新间隔，例如 500ms、2s、5s
```

示例：

```powershell
go run . --style classic
go run . --style overview
go run . --table
go run . --style overview --task --interval 1s
go run . --style classic --table --task --load
```

默认情况下，程序会保持更轻量的采样模式：

- 🧱 只显示上半区仪表盘
- 🚫 不显示进程表
- 🚫 不显示任务指标
- 🚫 不显示负载指标

## 🧭 项目结构

```text
.
├── .github
│   └── workflows
│       └── build.yml
├── internal
│   ├── dashboard
│   │   ├── helpers.go
│   │   ├── render.go
│   │   ├── types.go
│   │   └── render
│   │       ├── classic.go
│   │       └── overview.go
│   ├── metrics
│   │   ├── collect.go
│   │   ├── dashboard.go
│   │   ├── format.go
│   │   ├── paths_default.go
│   │   ├── paths_windows.go
│   │   ├── process.go
│   │   └── types.go
│   └── ui
│       ├── command.go
│       ├── model.go
│       ├── msg.go
│       ├── table.go
│       └── view.go
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## 🧠 internal 概览

### `internal/dashboard`

🎨 上半区仪表盘渲染核心。

- `types.go`：共享的渲染数据类型
- `render.go`：共享的绘制原语与通用样式
- `helpers.go`：宽度计算、格式化和通用辅助函数
- `render/classic.go`：经典双栏 htop 风格上半区
- `render/overview.go`：卡片式 overview 上半区

### `internal/metrics`

📡 实时系统采样与数据转换层。

- `collect.go`：主采集流程
- `dashboard.go`：将原始指标转换为 dashboard 渲染数据
- `process.go`：进程采样与排序逻辑
- `format.go`：共享输出格式化辅助函数
- `paths_default.go`：非 Windows 平台的磁盘根路径辅助函数
- `paths_windows.go`：Windows 平台的磁盘根路径辅助函数
- `types.go`：快照与指标结构体

### `internal/ui`

🫖 包裹渲染器的 Bubble Tea 应用外壳。

- `model.go`：应用状态与更新循环
- `view.go`：顶层界面拼装
- `table.go`：可选的下半区进程表
- `command.go`：异步刷新与加载动画命令
- `msg.go`：内部 Bubble Tea 消息类型

## 🤖 GitHub Actions

Workflow 文件：

- `.github/workflows/build.yml`
- [构建文档](.github/workflows/build.zh-cn.md)

版本来源：

- `internal/version/version.go` 是唯一真实来源
- workflow 会直接读取 `internal/version/version.go`
- release tag 会根据该文件自动创建为 `v<Number>`
- commit message 只控制 CI 是否执行，不负责传入版本号

行为：

- 🔁 Pull Request：始终构建产物
- 🧪 `build action`：只构建产物
- 🚀 `build release`：构建产物并创建 GitHub Release
- 🖱️ `workflow_dispatch`：手动触发构建

当前目标：

- Windows `amd64`
- Windows `arm64`
- Linux `amd64`
- Linux `arm64`
- macOS `amd64`
- macOS `arm64`
- Android `arm64`
- Android `amd64`
- Android `arm`

产物命名：

```text
htop-go-<target>-<version>
htop-go-<target>-<version>.exe
```

## 🛠️ 说明

- 🧪 本地缓存目录已被 `.gitignore` 忽略
- 📦 已包含 Android 构建，且其中部分 Android 目标会在 CI 中使用基于 NDK 的外部链接
- ⚡ `--table`、`--task` 和 `--load` 都是按需启用，因为它们比最小仪表盘路径更重
- 🧰 `internal/dashboard` 下的渲染层保持可复用，而 `internal/ui` 负责应用外壳
- 🔢 唯一版本来源是 `internal/version/version.go`

## 📦 技术栈

| 包 | 版本 | 作用 |
|:---|:---|:---|
| [![Go](https://img.shields.io/badge/Go-1.24.6-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/) | 1.24.6 | 编程语言与构建工具链 |
| [![bubbletea](https://img.shields.io/badge/bubbletea-1.3.4-111111?style=flat-square)](https://github.com/charmbracelet/bubbletea) | 1.3.4 | Bubble Tea TUI 运行时、事件循环与状态更新 |
| [![lipgloss](https://img.shields.io/badge/lipgloss-1.1.0-222222?style=flat-square)](https://github.com/charmbracelet/lipgloss) | 1.1.0 | 文本样式、布局、间距与彩色仪表盘渲染 |
| [![gopsutil](https://img.shields.io/badge/gopsutil-4.25.4-0F766E?style=flat-square)](https://github.com/shirou/gopsutil) | 4.25.4 | 跨平台 CPU、内存、交换区、进程与 uptime 采集 |
