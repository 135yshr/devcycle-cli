---
title: "features"
weight: 3
---

# features

Commands for managing DevCycle features (feature flags).

## list

List all features in a project.

### Usage

```bash
dvcx features list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List features in a specific project
$ dvcx features list -p my-app
KEY                 NAME                TYPE      STATUS
dark-mode           Dark Mode           release   active
new-checkout        New Checkout Flow   release   active
beta-features       Beta Features       permission inactive

# List features in JSON format
$ dvcx features list -p my-app -o json
[
  {
    "key": "dark-mode",
    "name": "Dark Mode",
    "type": "release",
    "status": "active"
  },
  ...
]

# Use default project from config
$ dvcx features list
```

### Notes

- If `--project` is not specified, the default project from configuration is used
- Feature types: `release`, `experiment`, `permission`, `ops`
- Feature status: `active`, `inactive`, `archived`

---

## get

Get details of a specific feature.

### Usage

```bash
dvcx features get <feature-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `feature-key` | The unique key of the feature |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get feature details
$ dvcx features get dark-mode -p my-app
KEY:         dark-mode
NAME:        Dark Mode
TYPE:        release
STATUS:      active
DESCRIPTION: Enable dark mode theme for the application
CREATED:     2024-01-20T10:00:00Z
UPDATED:     2024-06-15T14:30:00Z

# Get feature details in JSON format
$ dvcx features get dark-mode -p my-app -o json
{
  "key": "dark-mode",
  "name": "Dark Mode",
  "type": "release",
  "status": "active",
  "description": "Enable dark mode theme for the application",
  "createdAt": "2024-01-20T10:00:00Z",
  "updatedAt": "2024-06-15T14:30:00Z",
  "variables": [
    {
      "key": "dark-mode-enabled",
      "type": "Boolean"
    }
  ]
}
```

### Feature Types

| Type | Description |
|------|-------------|
| `release` | Standard feature flag for releasing features |
| `experiment` | A/B testing and experimentation |
| `permission` | User permission-based features |
| `ops` | Operational flags for system configuration |

### Notes

- The response includes associated variables and variations
- Use this command to verify feature configuration before deployment
