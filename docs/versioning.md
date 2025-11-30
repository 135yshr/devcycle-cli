# Versioning Guide

This document describes the versioning strategy for `dvcx`.

## Version Format

```text
X.Y.Z-HASH
```

| Component | Description | Example |
|-----------|-------------|---------|
| X | Major version (breaking changes) | `1` |
| Y | Minor version (new features) | `2` |
| Z | Patch version (bug fixes) | `3` |
| HASH | Git short commit hash (7 characters) | `abc1234` |

**Examples:**

- `1.0.0-abc1234` - Release version 1.0.0
- `dev-abc1234` - Development build (no tag)

## Semantic Versioning

We follow [Semantic Versioning 2.0.0](https://semver.org/):

- **MAJOR (X)**: Incompatible API changes
- **MINOR (Y)**: New functionality (backward compatible)
- **PATCH (Z)**: Bug fixes (backward compatible)

## Git Tags

Version tags follow the format:

```text
v*.*.* (e.g., v1.0.0, v1.2.3)
```

### Creating a Release

```bash
# Create and push a tag
git tag v1.0.0
git push origin v1.0.0
```

This triggers the release workflow which:

1. Builds binaries for all platforms
2. Creates a GitHub Release
3. Updates the Homebrew formula

## Build-time Version Injection

Version information is injected at build time using Go's `-ldflags`:

### Variables

| Variable | Description | Source |
|----------|-------------|--------|
| `cmd.Version` | Semantic version | Git tag (without `v` prefix) |
| `cmd.Commit` | Git short hash | `git rev-parse --short HEAD` |
| `cmd.Date` | Build timestamp | ISO 8601 format |

### Local Build

```bash
make build
./bin/dvcx version
```

Output:

```text
dvcx version 1.0.0-abc1234
  Commit:     abc1234
  Built:      2024-01-01T00:00:00Z
  Go version: go1.24
  OS/Arch:    darwin/arm64
```

### Development Build

Without a git tag:

```text
dvcx version dev-abc1234
  Commit:     abc1234
  Built:      2024-01-01T00:00:00Z
  Go version: go1.24
  OS/Arch:    darwin/arm64
```

## Implementation Details

### Source Code

Version variables are defined in `cmd/version.go`:

```go
var (
    Version = "dev"    // Set by ldflags
    Commit  = "none"   // Set by ldflags
    Date    = "unknown" // Set by ldflags
)

func GetVersion() string {
    return fmt.Sprintf("%s-%s", Version, Commit)
}
```

### Makefile

```makefile
VERSION=$(shell git describe --tags --abbrev=0 2>/dev/null | sed 's/^v//' || echo "dev")
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

LDFLAGS=-ldflags "\
    -s -w \
    -X github.com/135yshr/devcycle-cli/cmd.Version=$(VERSION) \
    -X github.com/135yshr/devcycle-cli/cmd.Commit=$(COMMIT) \
    -X github.com/135yshr/devcycle-cli/cmd.Date=$(DATE)"
```

### GoReleaser

In `.goreleaser.yaml`:

```yaml
builds:
  - ldflags:
      - -s -w
      - -X github.com/135yshr/devcycle-cli/cmd.Version={{.Version}}
      - -X github.com/135yshr/devcycle-cli/cmd.Commit={{.ShortCommit}}
      - -X github.com/135yshr/devcycle-cli/cmd.Date={{.Date}}
```

## Checking Version

### CLI

```bash
# Short version
dvcx --version

# Detailed version
dvcx version
```

### Make

```bash
make version
```

## Releasing

For release instructions, see [RELEASING.md](../RELEASING.md).
