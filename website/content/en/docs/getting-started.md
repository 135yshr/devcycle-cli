---
title: "Getting Started"
weight: 1
---

# Getting Started

This guide will help you install and set up dvcx on your system.

## Prerequisites

Before installing dvcx, you need:

1. **DevCycle Account**: Sign up at [devcycle.com](https://devcycle.com)
2. **API Credentials**: Obtain your Client ID and Client Secret from the DevCycle dashboard

### Getting API Credentials

1. Log in to your DevCycle dashboard
2. Navigate to **Settings** → **API Credentials**
3. Create new credentials or use existing ones
4. Note your **Client ID** and **Client Secret**

## Installation

### Using Go Install

If you have Go installed, the easiest way to install dvcx is:

```bash
go install github.com/135yshr/devcycle-cli@latest
```

### Download Binary

Download the latest binary for your platform from the [Releases page](https://github.com/135yshr/devcycle-cli/releases).

{{< tabs "installation" >}}
{{< tab "macOS (Apple Silicon)" >}}
{{< highlight bash >}}
curl -L https://github.com/135yshr/devcycle-cli/releases/latest/download/dvcx_darwin_arm64.tar.gz | tar xz
sudo mv dvcx /usr/local/bin/
{{< /highlight >}}
{{< /tab >}}
{{< tab "macOS (Intel)" >}}
{{< highlight bash >}}
curl -L https://github.com/135yshr/devcycle-cli/releases/latest/download/dvcx_darwin_amd64.tar.gz | tar xz
sudo mv dvcx /usr/local/bin/
{{< /highlight >}}
{{< /tab >}}
{{< tab "Linux" >}}
{{< highlight bash >}}
curl -L https://github.com/135yshr/devcycle-cli/releases/latest/download/dvcx_linux_amd64.tar.gz | tar xz
sudo mv dvcx /usr/local/bin/
{{< /highlight >}}
{{< /tab >}}
{{< tab "Windows" >}}
Download the `.zip` file from the releases page and add the extracted folder to your PATH.
{{< /tab >}}
{{< /tabs >}}

### Build from Source

```bash
git clone https://github.com/135yshr/devcycle-cli.git
cd devcycle-cli
make build
```

The binary will be created at `bin/dvcx`.

## Authentication

After installation, authenticate with your DevCycle credentials:

```bash
dvcx auth login
```

You will be prompted to enter:

- **Client ID**: Your DevCycle API client ID
- **Client Secret**: Your DevCycle API client secret

The credentials are securely stored in `.devcycle/token.json` in your project directory.

## Verify Installation

Check that dvcx is installed correctly:

```bash
dvcx version
```

This should display version information:

```
dvcx version v0.1.0
  commit: abc1234
  built at: 2025-01-15T10:00:00Z
  go version: go1.24
  platform: darwin/arm64
```

## First Commands

Now that you're set up, try these commands:

```bash
# List all your projects
dvcx projects list

# Get details of a specific project
dvcx projects get my-project-key

# List features in a project
dvcx features list -p my-project-key
```

## Next Steps

- Learn about all available [Commands]({{< relref "/docs/commands" >}})
- Configure [default settings]({{< relref "/docs/configuration" >}})
- Read the [FAQ]({{< relref "/docs/faq" >}}) for common questions
