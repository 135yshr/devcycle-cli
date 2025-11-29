---
title: "version"
weight: 10
---

# version

Display version information about dvcx.

## Usage

```bash
dvcx version
```

## Output

The command displays:

- **Version**: The semantic version number
- **Commit**: Git commit hash of the build
- **Built at**: Build timestamp
- **Go version**: Go compiler version used
- **Platform**: Operating system and architecture

## Example

```bash
$ dvcx version
dvcx version v0.1.0
  commit: abc1234
  built at: 2025-01-15T10:00:00Z
  go version: go1.24
  platform: darwin/arm64
```

## Short Version

You can also use the `--version` flag on the root command:

```bash
$ dvcx --version
dvcx version v0.1.0
```

## Notes

- Version information is injected at build time using Go's ldflags
- The commit hash helps identify the exact source code version
- Useful for debugging and reporting issues
