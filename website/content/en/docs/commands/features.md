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

---

## create

Create a new feature in a project.

### Usage

```bash
# Simple create (v1 API)
dvcx features create --name <name> --key <key> [flags]

# Create from JSON file (v2 API)
dvcx features create --from-file <file.json> [flags]

# Create from stdin (v2 API)
dvcx features create --from-file - [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--name` | `-n` | Feature name | Yes (for simple create) |
| `--key` | `-k` | Feature key | Yes (for simple create) |
| `--description` | `-d` | Feature description | No |
| `--type` | `-t` | Feature type (release, experiment, permission, ops) | No (default: release) |
| `--from-file` | `-F` | JSON input file for feature creation (uses v2 API), use `-` for stdin | No |
| `--dry-run` | | Validate configuration without creating | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Simple Create Example

```bash
# Create a release feature
$ dvcx features create -p my-app -n "Dark Mode" -k dark-mode
KEY:         dark-mode
NAME:        Dark Mode
TYPE:        release
STATUS:      active
CREATED:     2024-06-20T10:00:00Z

# Create an experiment feature with description
$ dvcx features create -p my-app -n "New Checkout" -k new-checkout -t experiment -d "A/B test for new checkout flow"

# Create feature and output as JSON
$ dvcx features create -p my-app -n "Beta Feature" -k beta-feature -o json
```

### Create from JSON File (v2 API)

The `--from-file` flag enables full feature configuration using the v2 API, including variables, variations, and targeting rules.

```bash
# Create from JSON file
$ dvcx features create -p my-app --from-file feature.json

# Validate without creating (dry-run)
$ dvcx features create -p my-app --from-file feature.json --dry-run
```

#### JSON File Format

```json
{
  "name": "Dark Mode",
  "key": "dark-mode",
  "description": "Enable dark mode theme",
  "type": "release",
  "tags": ["ui", "theme"],
  "variables": [
    {
      "key": "enabled",
      "name": "Enabled",
      "type": "Boolean"
    }
  ],
  "variations": [
    {
      "key": "off",
      "name": "Off",
      "variables": { "enabled": false }
    },
    {
      "key": "on",
      "name": "On",
      "variables": { "enabled": true }
    }
  ],
  "controlVariation": "off",
  "configurations": {
    "development": {
      "status": "active",
      "targets": [
        {
          "name": "All Users",
          "audience": {
            "filters": {
              "operator": "and",
              "filters": [{ "type": "all" }]
            }
          },
          "distribution": [
            { "_variation": "on", "percentage": 1.0 }
          ]
        }
      ]
    }
  }
}
```

### Create from stdin

You can pipe JSON content directly to the command using `-` as the file path.

```bash
# Pipe from file
$ cat feature.json | dvcx features create -p my-app --from-file -

# Use heredoc
$ dvcx features create -p my-app --from-file - <<EOF
{
  "name": "Quick Feature",
  "key": "quick-feature",
  "type": "release"
}
EOF

# Pipe from another command (e.g., jq)
$ jq '.features[0]' features.json | dvcx features create -p my-app --from-file -
```

### v2 API Supported Fields

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Feature name (required) |
| `key` | string | Feature key (required) |
| `description` | string | Feature description |
| `type` | string | Feature type: release, experiment, permission, ops |
| `tags` | string[] | Tags for organizing features |
| `variables` | object[] | Variable definitions |
| `variations` | object[] | Variation definitions with variable values |
| `controlVariation` | string | Key of the control variation |
| `configurations` | object | Environment-specific targeting configurations |
| `sdkVisibility` | object | SDK visibility settings (mobile, client, server) |
| `settings` | object | Feature settings (publicName, publicDescription, optInEnabled) |

### Notes

- Feature keys must be unique within a project
- Feature keys can contain lowercase letters, numbers, hyphens, and underscores
- Default feature type is `release` if not specified
- Use `--from-file -` to read JSON from stdin
- Use `--dry-run` to validate configuration without creating the feature

---

## update

Update an existing feature.

### Usage

```bash
dvcx features update <feature-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `feature-key` | The unique key of the feature to update |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--name` | `-n` | New feature name | No |
| `--description` | `-d` | New feature description | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Update feature name
$ dvcx features update dark-mode -p my-app -n "Dark Theme"

# Update feature description
$ dvcx features update dark-mode -p my-app -d "Enable dark theme for the application"

# Update both name and description
$ dvcx features update dark-mode -p my-app -n "Dark Theme" -d "Enable dark theme"
```

### Notes

- Only the specified fields will be updated
- Feature key and type cannot be changed after creation

---

## delete

Delete a feature from a project.

### Usage

```bash
dvcx features delete <feature-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `feature-key` | The unique key of the feature to delete |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--force` | `-f` | Skip confirmation prompt | No |

### Example

```bash
# Delete a feature (with confirmation prompt)
$ dvcx features delete dark-mode -p my-app
Are you sure you want to delete feature 'dark-mode'? [y/N]: y
Feature 'dark-mode' deleted successfully

# Delete a feature without confirmation
$ dvcx features delete dark-mode -p my-app --force
Feature 'dark-mode' deleted successfully
```

### Notes

- Deleting a feature will also delete all associated variables and targeting rules
- This action cannot be undone
- Use `--force` flag to skip confirmation in automated scripts
