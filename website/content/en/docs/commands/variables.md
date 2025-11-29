---
title: "variables"
weight: 4
---

# variables

Commands for managing DevCycle variables.

Variables are the values that feature flags control. Each feature can have multiple variables of different types.

## list

List all variables in a project.

### Usage

```bash
dvcx variables list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List variables in a project
$ dvcx variables list -p my-app
KEY                   TYPE      FEATURE
dark-mode-enabled     Boolean   dark-mode
theme-color           String    dark-mode
checkout-version      Number    new-checkout
config-json           JSON      beta-features

# List variables in JSON format
$ dvcx variables list -p my-app -o json
[
  {
    "key": "dark-mode-enabled",
    "type": "Boolean",
    "feature": "dark-mode"
  },
  ...
]
```

### Variable Types

| Type | Description | Example Values |
|------|-------------|----------------|
| `Boolean` | True/false values | `true`, `false` |
| `String` | Text values | `"dark"`, `"light"` |
| `Number` | Numeric values | `1`, `2.5`, `100` |
| `JSON` | Complex JSON objects | `{"key": "value"}` |

---

## get

Get details of a specific variable.

### Usage

```bash
dvcx variables get <variable-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `variable-key` | The unique key of the variable |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get variable details
$ dvcx variables get dark-mode-enabled -p my-app
KEY:          dark-mode-enabled
TYPE:         Boolean
FEATURE:      dark-mode
DESCRIPTION:  Controls whether dark mode is enabled
CREATED:      2024-01-20T10:00:00Z

# Get variable details in JSON format
$ dvcx variables get dark-mode-enabled -p my-app -o json
{
  "key": "dark-mode-enabled",
  "type": "Boolean",
  "feature": "dark-mode",
  "description": "Controls whether dark mode is enabled",
  "createdAt": "2024-01-20T10:00:00Z",
  "updatedAt": "2024-06-15T14:30:00Z"
}
```

### Notes

- Variables are always associated with a feature
- The variable key must be unique within a project
- Variable types cannot be changed after creation

---

## create

Create a new variable in a project.

### Usage

```bash
dvcx variables create [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--name` | `-n` | Variable name | Yes |
| `--key` | `-k` | Variable key | Yes |
| `--description` | `-d` | Variable description | No |
| `--type` | `-t` | Variable type (String, Boolean, Number, JSON) | No (default: Boolean) |
| `--feature` | | Associated feature key | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Create a boolean variable
$ dvcx variables create -p my-app -n "Dark Mode Enabled" -k dark-mode-enabled
KEY:          dark-mode-enabled
TYPE:         Boolean
STATUS:       active
CREATED:      2024-06-20T10:00:00Z

# Create a string variable with description
$ dvcx variables create -p my-app -n "Theme Color" -k theme-color -t String -d "Primary theme color"

# Create a variable associated with a feature
$ dvcx variables create -p my-app -n "Checkout Version" -k checkout-version -t Number --feature new-checkout

# Create variable and output as JSON
$ dvcx variables create -p my-app -n "Config" -k config-json -t JSON -o json
```

### Notes

- Variable keys must be unique within a project
- Variable keys can contain lowercase letters, numbers, hyphens, and underscores
- Default variable type is `Boolean` if not specified
- Use `--feature` to associate the variable with an existing feature

---

## update

Update an existing variable.

### Usage

```bash
dvcx variables update <variable-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `variable-key` | The unique key of the variable to update |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--name` | `-n` | New variable name | No |
| `--description` | `-d` | New variable description | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Update variable name
$ dvcx variables update dark-mode-enabled -p my-app -n "Dark Theme Enabled"

# Update variable description
$ dvcx variables update dark-mode-enabled -p my-app -d "Controls whether dark theme is enabled"

# Update both name and description
$ dvcx variables update dark-mode-enabled -p my-app -n "Dark Theme" -d "Dark theme toggle"
```

### Notes

- Only the specified fields will be updated
- Variable key and type cannot be changed after creation

---

## delete

Delete a variable from a project.

### Usage

```bash
dvcx variables delete <variable-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `variable-key` | The unique key of the variable to delete |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--force` | `-f` | Skip confirmation prompt | No |

### Example

```bash
# Delete a variable (with confirmation prompt)
$ dvcx variables delete dark-mode-enabled -p my-app
Are you sure you want to delete variable 'dark-mode-enabled'? [y/N]: y
Variable 'dark-mode-enabled' deleted successfully

# Delete a variable without confirmation
$ dvcx variables delete dark-mode-enabled -p my-app --force
Variable 'dark-mode-enabled' deleted successfully
```

### Notes

- This action cannot be undone
- Use `--force` flag to skip confirmation in automated scripts
- Variables must be disassociated from features before deletion in some cases
