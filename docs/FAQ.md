# Frequently Asked Questions

## General

### What is dvcx?

`dvcx` (DevCycle CLI Extended) is an unofficial command-line tool for the
[DevCycle Management API](https://docs.devcycle.com/management-api/).
It provides comprehensive access to DevCycle's feature flag management capabilities
from the command line.

### How is dvcx different from the official DevCycle CLI?

The [official DevCycle CLI](https://docs.devcycle.com/cli/) focuses on SDK code generation
and basic operations. `dvcx` aims to provide:

- Full CRUD operations for all Management API resources
- Bulk operations and automation capabilities
- Detailed audit log access
- Advanced targeting and audience management

### Is dvcx affiliated with DevCycle?

No, `dvcx` is an unofficial, community-driven project. It is not affiliated with
or endorsed by DevCycle. Use at your own risk.

## Installation

### How do I install dvcx?

**Using Homebrew (recommended):**

```bash
brew tap 135yshr/tap
brew install dvcx
```

**From source:**

```bash
git clone https://github.com/135yshr/devcycle-cli.git
cd devcycle-cli
make build
```

### What are the system requirements?

- Go 1.24+ (for building from source)
- macOS, Linux, or Windows

## Configuration

### Where should I store my API credentials?

Create a `.devcycle/config.yaml` file in your project root:

```yaml
client_id: your-client-id
client_secret: your-client-secret
project: your-project-key
```

**Important:** Never commit credentials to version control.
The `.devcycle/` directory is already included in the default `.gitignore`.

### Can I use environment variables instead of a config file?

Yes, you can use environment variables with the `DVCX_` prefix:

```bash
export DVCX_CLIENT_ID=your-client-id
export DVCX_CLIENT_SECRET=your-client-secret
export DVCX_PROJECT=your-project-key
```

### How do I get DevCycle API credentials?

1. Log in to the [DevCycle Dashboard](https://app.devcycle.com/)
2. Navigate to Settings > API Credentials
3. Create a new API Key with appropriate permissions

## Troubleshooting

### I'm getting authentication errors

- Verify your `client_id` and `client_secret` are correct
- Check that your API credentials have the necessary permissions
- Ensure your credentials haven't expired

### Commands are returning empty results

- Verify you have the correct `project` key configured
- Check that you have access to the project in DevCycle
- Try running with `DVCX_DEBUG=true` for verbose output

### How do I enable debug mode?

Set the environment variable:

```bash
DVCX_DEBUG=true dvcx projects list
```

Or add to your config file:

```yaml
debug: true
```

## Contributing

### How can I contribute to dvcx?

See our [Contributing Guide](contributing.md) for details on:

- Reporting issues
- Submitting pull requests
- Development workflow
- Code style guidelines

### I found a security vulnerability. How do I report it?

Please see our [Security Policy](../SECURITY.md) for responsible disclosure guidelines.
Do not report security vulnerabilities through public GitHub issues.

## Support

### Where can I get help?

- [GitHub Issues](https://github.com/135yshr/devcycle-cli/issues) - Bug reports and feature requests
- [GitHub Discussions](https://github.com/135yshr/devcycle-cli/discussions) - Questions and discussions

### Is there commercial support available?

No, this is a community project without commercial support.
For official DevCycle support, please contact [DevCycle](https://devcycle.com/) directly.
