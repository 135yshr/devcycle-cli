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
