---
title: "FAQ"
weight: 4
---

# Frequently Asked Questions

## Installation

### How do I install dvcx?

The easiest way is using Go:

```bash
go install github.com/135yshr/devcycle-cli@latest
```

Or download a binary from the [Releases page](https://github.com/135yshr/devcycle-cli/releases).

### Where can I find my DevCycle API credentials?

1. Log in to your [DevCycle dashboard](https://app.devcycle.com)
2. Go to **Settings** → **API Credentials**
3. Create new credentials or copy existing ones

## Authentication

### My authentication token expired. What should I do?

Simply run `dvcx auth login` again to get a new token.

### Where is my token stored?

The token is stored in `.devcycle/token.json` in your current project directory.

### How do I use dvcx in CI/CD?

Set the `DVCX_CLIENT_ID` and `DVCX_CLIENT_SECRET` environment variables:

```bash
export DVCX_CLIENT_ID=your-client-id
export DVCX_CLIENT_SECRET=your-client-secret
dvcx auth login
```

## Commands

### How do I set a default project?

Create a `.devcycle/config.yaml` file:

```yaml
project: your-project-key
```

### How do I get JSON output?

Use the `-o json` flag:

```bash
dvcx projects list -o json
```

### What's the difference between dvcx and the official DevCycle CLI?

dvcx is an unofficial tool that provides:

- Full Management API access
- Multiple output formats (table, JSON, YAML)
- Project-scoped configuration
- Additional features not available in the official CLI

## Troubleshooting

### "Project key is required" error

Either:

1. Specify the project with `-p` flag: `dvcx features list -p my-project`
2. Set a default project in `.devcycle/config.yaml`

### "Unauthorized" error

Your token may have expired. Run `dvcx auth login` to re-authenticate.

### Command not found

Make sure the binary is in your PATH:

```bash
# Check if dvcx is accessible
which dvcx

# If using go install, ensure GOPATH/bin is in PATH
export PATH=$PATH:$(go env GOPATH)/bin
```

## Contributing

### How can I contribute?

See the [Contributing Guide](https://github.com/135yshr/devcycle-cli/blob/main/docs/contributing.md) for details on:

- Setting up the development environment
- Submitting pull requests
- Coding standards

### Where do I report bugs?

Open an issue on [GitHub Issues](https://github.com/135yshr/devcycle-cli/issues).
