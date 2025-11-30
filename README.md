# dvcx - DevCycle CLI Extended

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)](https://go.dev/)

An unofficial command-line tool for the [DevCycle Management API](https://docs.devcycle.com/management-api/).

[日本語ドキュメント](docs/ja/README.md)

## Why dvcx?

While DevCycle provides an [official CLI](https://docs.devcycle.com/cli/), it has limited functionality for advanced operations. `dvcx` aims to provide comprehensive access to the Management API, enabling:

- Full CRUD operations for all resources (Projects, Features, Variables, etc.)
- Bulk operations and automation
- Detailed audit log access
- Advanced targeting and audience management

## Installation

### From Source

```bash
git clone https://github.com/135yshr/devcycle-cli.git
cd devcycle-cli
make build
```

The binary will be available at `bin/dvcx`.

### Go Install

```bash
go install github.com/135yshr/devcycle-cli@latest
```

## Quick Start

1. **Configure credentials**

   Create `.devcycle/config.yaml` in your project root:

   ```yaml
   client_id: your-client-id
   client_secret: your-client-secret
   project: your-project-key
   ```

2. **Login**

   ```bash
   dvcx auth login
   ```

3. **List projects**

   ```bash
   dvcx projects list
   ```

## Commands

> Note: This project is under active development. See [Roadmap](docs/roadmap.md) for implementation status.

### Core Commands

| Command | Description | Status |
|---------|-------------|--------|
| `auth login/logout` | Authentication | ✅ Completed |
| `projects list/get/create/update` | Project management | ✅ Completed |
| `features list/get/create/update/delete` | Feature flags | ✅ Completed |
| `variables list/get/create/update/delete` | Variables | ✅ Completed |
| `environments list/get` | Environments | ✅ Completed |
| `targeting get/update/enable/disable` | Feature targeting | ✅ Completed |
| `variations list/get/create/update/delete` | Feature variations | ✅ Completed |

### Audiences & Overrides (Phase 4)

| Command | Description | Status |
|---------|-------------|--------|
| `audiences list/get/create/update/delete` | Audience management | ✅ Completed |
| `overrides list/get/set/delete` | Self-targeting overrides | ✅ Completed |

### Operations & Monitoring (Phase 5)

| Command | Description | Status |
|---------|-------------|--------|
| `audit list/feature` | Audit logs | ✅ Completed |
| `metrics list/get/create/update/delete/results` | Metrics management | ✅ Completed |
| `webhooks list/get/create/update/delete` | Webhook management | ✅ Completed |
| `custom-properties list/get/create/update/delete` | Custom properties | ✅ Completed |

See [docs/roadmap.md](docs/roadmap.md) for the complete list and [docs/usage.md](docs/usage.md) for detailed usage examples.

## Documentation

- [Usage Guide](docs/usage.md) - Detailed command usage examples
- [API Reference](docs/api-reference.md) - DevCycle Management API endpoints
- [Roadmap](docs/roadmap.md) - Implementation phases and progress
- [Contributing](docs/contributing.md) - How to contribute
- [Development](docs/development.md) - Development setup and architecture

## Development

### Prerequisites

- Go 1.24+
- [pre-commit](https://pre-commit.com/) - For Git hooks
- [markdownlint-cli2](https://github.com/DavidAnson/markdownlint-cli2) - For Markdown linting

### Setup

```bash
# Install pre-commit and markdownlint-cli2
brew install pre-commit markdownlint-cli2

# Install pre-commit hooks
pre-commit install
```

See [docs/development.md](docs/development.md) for detailed setup instructions.

## Contributing

Contributions are welcome! Please read our [Contributing Guide](docs/contributing.md) before submitting a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Disclaimer

This is an unofficial tool and is not affiliated with or endorsed by DevCycle. Use at your own risk.
