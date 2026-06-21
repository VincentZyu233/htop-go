# htop-go

🖥️📊⚡ A cross-platform `htop`-style TUI built with Go, Bubble Tea, Lip Gloss, and gopsutil.

> **[📖 English](README.md)**
> **[📖 简体中文](README.zh-cn.md)**

[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/VincentZyuApps/htop-go)
[![Gitee](https://img.shields.io/badge/Gitee-C71D23?style=for-the-badge&logo=gitee&logoColor=white)](https://gitee.com/vincent-zyu/htop-go)

[![Go 1.24.6](https://img.shields.io/badge/Go-1.24.6-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

[![CI](https://img.shields.io/github/actions/workflow/status/VincentZyuApps/htop-go/build.yml?style=for-the-badge&logo=githubactions&logoColor=white&label=CI)](https://github.com/VincentZyuApps/htop-go/actions/workflows/build.yml)

## 🌍 Platform Goals

- 🪟 Windows
- 🐧 Linux
- 🍎 macOS
- 🤖 Android / Termux

## ✨ Features

- 🎛️ Two upper-panel styles: `classic` and `overview`
- 📋 Optional lower process table via `--table`
- 🧩 Stateless dashboard render layer under `internal/dashboard`
- 📱 Cross-build friendly for Windows, Linux, macOS, and Android / Termux
- 🚀 GitHub Actions build and release flow controlled by commit message

## 🚀 Run Locally

```powershell
$env:GOCACHE = "$PWD\.gocache"
$env:GOMODCACHE = "$PWD\.gomodcache"
$env:GOPATH = "$PWD\.gopath"
go mod tidy
go run .
```

Quick checks:

```powershell
go run . --help
go run . --style classic
go run . --style overview
go run . --table --task
```

## 📥 Install

One-liner install script for Debian / Ubuntu, Fedora / RHEL, and Termux:

```bash
curl -fsSL https://raw.githubusercontent.com/VincentZyuApps/htop-go/main/scripts/install.sh | bash
htop-go --help
```

China-friendly Gitee mirror:

```bash
curl -fsSL https://gitee.com/vincent-zyu/htop-go/raw/main/scripts/install_gitee.sh | bash
htop-go --help
```

Manual version pin:

```bash
HTOP_GO_VERSION=v0.1.0 bash -c "$(curl -fsSL https://raw.githubusercontent.com/VincentZyuApps/htop-go/main/scripts/install.sh)"
```

> ⚠️ The two `curl ... | bash` install scripts support `apt`, `dnf`, and Termux. They install the release binary directly into `PATH`.

## ⚙️ CLI

```text
--style classic|overview   choose upper dashboard style
--table                    enable lower process table and process sampling
--task                     enable task metrics
--load                     enable load metrics
-t, --interval             refresh interval, e.g. 500ms, 2s, 5s
```

Examples:

```powershell
go run . --style classic
go run . --style overview
go run . --table
go run . --style overview --task --interval 1s
go run . --style classic --table --task --load
```

By default, the app keeps sampling lighter:

- 🧱 upper dashboard only
- 🚫 no process table
- 🚫 no task metrics
- 🚫 no load metrics

## 🧭 Project Layout

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

## 🧠 Internal Overview

### `internal/dashboard`

🎨 Upper dashboard render core.

- `types.go`: shared render data types
- `render.go`: shared drawing primitives and common styles
- `helpers.go`: width math, formatting, and shared helpers
- `render/classic.go`: classic split-column htop-like upper panel
- `render/overview.go`: overview card-style upper panel

### `internal/metrics`

📡 Live system sampling and conversion layer.

- `collect.go`: main metrics collection flow
- `dashboard.go`: converts raw metrics to dashboard-facing data
- `process.go`: process sampling and sorting
- `format.go`: shared output helpers
- `paths_default.go`: non-Windows disk root helpers
- `paths_windows.go`: Windows disk root helpers
- `types.go`: snapshot and metric structs

### `internal/ui`

🫖 Bubble Tea app shell around the renderer.

- `model.go`: app state and update loop
- `view.go`: top-level screen composition
- `table.go`: optional lower process table
- `command.go`: async refresh and spinner commands
- `msg.go`: internal Bubble Tea messages

## 🤖 GitHub Actions

Workflow file:

- `.github/workflows/build.yml`
- [Build Docs](.github/workflows/build.md)

Version source:

- `internal/version/version.go` is the single source of truth
- the workflow reads `internal/version/version.go` directly
- the release tag is created automatically as `v<Number>` from that file
- build/release commit messages only control whether CI runs, not the version number

Behavior:

- 🔁 Pull requests: always build artifacts
- 🧪 `build action`: build artifacts only
- 🚀 `build release`: build artifacts and create a GitHub Release
- 🖱️ `workflow_dispatch`: manual build

Current targets:

- Windows `amd64`
- Windows `arm64`
- Linux `amd64`
- Linux `arm64`
- macOS `amd64`
- macOS `arm64`
- Android `arm64`
- Android `amd64`
- Android `arm`

Artifact naming:

```text
htop-go-<target>-<version>
htop-go-<target>-<version>.exe
```

## 🛠️ Notes

- 🧪 Local cache directories are ignored by `.gitignore`
- 📦 Android builds are included, and some Android targets use NDK-based external linking in CI
- ⚡ `--table`, `--task`, and `--load` are opt-in because they are heavier than the minimal dashboard path
- 🧰 The render layer under `internal/dashboard` stays reusable, while `internal/ui` is the app shell
- 🔢 The only version source is `internal/version/version.go`

## 📦 Tech Stack

| Package | Version | Role |
|:---|:---|:---|
| [![Go](https://img.shields.io/badge/Go-1.24.6-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev/) | 1.24.6 | Programming language and toolchain |
| [![bubbletea](https://img.shields.io/badge/bubbletea-1.3.4-111111?style=flat-square)](https://github.com/charmbracelet/bubbletea) | 1.3.4 | Bubble Tea TUI runtime, event loop, and state updates |
| [![lipgloss](https://img.shields.io/badge/lipgloss-1.1.0-222222?style=flat-square)](https://github.com/charmbracelet/lipgloss) | 1.1.0 | Text styling, layout, spacing, and colored dashboard rendering |
| [![gopsutil](https://img.shields.io/badge/gopsutil-4.25.4-0F766E?style=flat-square)](https://github.com/shirou/gopsutil) | 4.25.4 | Cross-platform CPU, memory, swap, process, and uptime collection |
