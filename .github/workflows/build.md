# Build & Release Workflow

> **[📖 English](build.md)**
> **[📖 简体中文](build.zh-cn.md)**

## 📋 Overview

This workflow is intentionally simple.

- version comes only from `internal/version/version.go`
- CI reads that file and turns `0.1.0` into release tag `v0.1.0`
- commit message keywords only control whether CI builds or releases

## 🔑 Keywords

| Keyword in commit message | Build artifacts | GitHub Release |
|---------------------------|:---:|:---:|
| `build action` | ✅ | ❌ |
| `build release` | ✅ | ✅ |

Notes:

- Pull requests always build
- `workflow_dispatch` also builds
- if no keyword is present on `main`, CI skips build and release

## 🚀 Usage Examples

```bash
# Build only
git commit --allow-empty -m "ci: verify matrix (build action)"

# Build and create GitHub Release
git commit --allow-empty -m "release: ship 0.1.0 (build release)"

# Normal commit, no CI build/release
git commit -m "docs: update README"
```

## 🏗️ Build Targets

| Platform | Architecture | Notes |
|----------|:---:|-------|
| Windows | `amd64` | `.exe` binary |
| Windows | `arm64` | `.exe` binary |
| Linux | `amd64` | native Go binary |
| Linux | `arm64` | native Go binary |
| macOS | `amd64` | native Go binary |
| macOS | `arm64` | native Go binary |
| Android | `arm64` | Termux-friendly target |
| Android | `amd64` | emulator / x86_64 Android target |
| Android | `arm` | 32-bit ARM Android target |

## 📦 Pipeline Stages

```text
check -> build -> release
```

### `check`

- parses commit message
- reads `internal/version/version.go`
- exports `v<Number>` as workflow version

### `build`

- uses Go from `go.mod`
- cross-builds all configured targets
- uploads artifacts

### `release`

- downloads artifacts
- renders `.github/release_template.md`
- creates GitHub Release with tag `v<Number>`

## 📌 Version Rule

The single version source is:

```text
internal/version/version.go
```

Current form:

```go
const Number = "0.1.0"
```

When you want a new release:

1. update `internal/version/version.go`
2. commit the change
3. use `build release` in the commit message

## 🔗 Related Files

- `README.md`
- `README.zh-cn.md`
- `.github/workflows/build.yml`
- `.github/release_template.md`
- `internal/version/version.go`
