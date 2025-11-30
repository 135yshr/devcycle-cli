---
title: "overrides"
weight: 11
---

# overrides

Commands for managing self-targeting overrides. Overrides allow developers to set specific variations for themselves during development and testing.

## list

List all overrides for a feature.

### Usage

```bash
dvcx overrides list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List all overrides for a feature
$ dvcx overrides list -p my-app -f dark-mode
USER ID          ENVIRONMENT    VARIATION
user-123         development    on
user-456         development    off
user-789         staging        on

# List overrides in JSON format
$ dvcx overrides list -p my-app -f dark-mode -o json
```

---

## get

Get your current override for a feature in a specific environment.

### Usage

```bash
dvcx overrides get [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--environment` | `-e` | Environment key | Yes |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get your override for a feature
$ dvcx overrides get -p my-app -f dark-mode -e development
FEATURE      ENVIRONMENT    VARIATION
dark-mode    development    on

# Get override in JSON format
$ dvcx overrides get -p my-app -f dark-mode -e development -o json
{
  "feature": "dark-mode",
  "environment": "development",
  "variation": "on"
}
```

---

## set

Set an override for yourself on a specific feature and environment.

### Usage

```bash
dvcx overrides set [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--environment` | `-e` | Environment key | Yes |
| `--variation` | `-v` | Variation key to serve | Yes |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Set an override to enable a feature for yourself
$ dvcx overrides set -p my-app -f dark-mode -e development --variation on
Override set successfully

# Set override to disable a feature
$ dvcx overrides set -p my-app -f dark-mode -e development --variation off
```

### Notes

- Overrides only affect your own user ID
- Overrides are per-environment
- The variation key must be a valid variation for the feature

---

## delete

Delete your override for a specific feature and environment.

### Usage

```bash
dvcx overrides delete [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--environment` | `-e` | Environment key | Yes |

### Example

```bash
# Delete your override for a feature
$ dvcx overrides delete -p my-app -f dark-mode -e development
Override deleted successfully
```

---

## list-mine

List all your overrides across all features in a project.

### Usage

```bash
dvcx overrides list-mine [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List all your overrides
$ dvcx overrides list-mine -p my-app
FEATURE        ENVIRONMENT    VARIATION
dark-mode      development    on
new-checkout   development    variation-b
beta-feature   staging        enabled

# List in JSON format
$ dvcx overrides list-mine -p my-app -o json
```

---

## delete-mine

Delete all your overrides in a project.

### Usage

```bash
dvcx overrides delete-mine [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--force` | | Skip confirmation prompt | No |

### Example

```bash
# Delete all your overrides (with confirmation)
$ dvcx overrides delete-mine -p my-app
Are you sure you want to delete all your overrides? [y/N]: y
All overrides deleted successfully

# Delete without confirmation
$ dvcx overrides delete-mine -p my-app --force
```

### Notes

- This deletes all your overrides across all features and environments
- This action cannot be undone
- Useful for cleaning up after testing
