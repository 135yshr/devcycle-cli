---
title: "Configuration"
weight: 3
---

# Configuration

dvcx uses a YAML configuration file to store settings and defaults.

## Configuration File Location

The configuration file is located at `.devcycle/config.yaml` in your project directory.

```
your-project/
├── .devcycle/
│   ├── config.yaml    # Configuration file
│   └── token.json     # Authentication token (auto-generated)
├── src/
└── ...
```

## Configuration Options

### Example Configuration

```yaml
# .devcycle/config.yaml

# Default project key
project: my-app

# Default output format (table, json, yaml)
output: table
```

### Available Options

| Option | Type | Description | Default |
|--------|------|-------------|---------|
| `project` | string | Default project key for commands | (none) |
| `output` | string | Default output format | `table` |

## Setting Default Project

Instead of specifying `--project` for every command, set a default:

```yaml
# .devcycle/config.yaml
project: my-app
```

Now you can run commands without the `-p` flag:

```bash
# Before: required project flag
dvcx features list -p my-app

# After: uses default from config
dvcx features list
```

## Output Format

Set your preferred default output format:

```yaml
output: json  # or table, yaml
```

## Environment Variables

Configuration options can also be set via environment variables with the `DVCX_` prefix:

| Environment Variable | Configuration Option |
|---------------------|---------------------|
| `DVCX_PROJECT` | `project` |
| `DVCX_OUTPUT` | `output` |

### Example

```bash
export DVCX_PROJECT=my-app
export DVCX_OUTPUT=json

dvcx features list  # Uses environment variables
```

## Priority Order

Configuration values are resolved in the following order (highest priority first):

1. **Command-line flags** (`--project`, `--output`)
2. **Environment variables** (`DVCX_PROJECT`, `DVCX_OUTPUT`)
3. **Configuration file** (`.devcycle/config.yaml`)
4. **Default values**

## Authentication Token

The authentication token is stored separately in `.devcycle/token.json`:

```json
{
  "access_token": "...",
  "expires_at": "2025-01-15T12:00:00Z"
}
```

{{< hint warning >}}
**Security Note**: Add `.devcycle/token.json` to your `.gitignore` to prevent accidentally committing credentials.
{{< /hint >}}

### Recommended .gitignore

```gitignore
# DevCycle CLI
.devcycle/token.json
```

## Custom Config File Path

Use the `--config` flag to specify a custom configuration file:

```bash
dvcx --config /path/to/custom-config.yaml projects list
```
