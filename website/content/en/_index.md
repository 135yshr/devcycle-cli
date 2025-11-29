---
title: "dvcx - DevCycle CLI Extended"
type: docs
---

# dvcx

**Unofficial CLI tool for DevCycle Management API**

dvcx provides command-line access to DevCycle features that are not available or limited in the official CLI.

## Features

- **Full Management API Access**: Access all DevCycle Management API endpoints
- **Multiple Output Formats**: Table, JSON, and YAML output formats
- **Project-scoped Configuration**: Store default project settings per directory
- **Cross-platform**: Works on macOS, Linux, and Windows

## Quick Start

```bash
# Install
go install github.com/135yshr/devcycle-cli@latest

# Login with your DevCycle credentials
dvcx auth login

# List your projects
dvcx projects list

# List features in a project
dvcx features list -p your-project-key
```

## Documentation

{{< columns >}}

### [Getting Started]({{< relref "/docs/getting-started" >}})

Learn how to install and configure dvcx.

<--->

### [Commands]({{< relref "/docs/commands" >}})

Complete reference for all CLI commands.

<--->

### [Configuration]({{< relref "/docs/configuration" >}})

Learn about configuration options.

{{< /columns >}}

## GitHub

- [Repository](https://github.com/135yshr/devcycle-cli)
- [Issues](https://github.com/135yshr/devcycle-cli/issues)
- [Releases](https://github.com/135yshr/devcycle-cli/releases)
