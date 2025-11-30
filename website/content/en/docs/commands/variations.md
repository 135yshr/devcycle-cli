---
title: "variations"
weight: 7
---

# variations

Commands for managing DevCycle feature variations.

Variations define the different values that a feature can serve to users. Each variation contains a set of variable values that are delivered together.

## list

List all variations for a feature.

### Usage

```bash
dvcx variations list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List variations for a feature
$ dvcx variations list -p my-app -f dark-mode
KEY           NAME           VARIABLES
variation-on  Dark Mode On   {"enabled": true}
variation-off Dark Mode Off  {"enabled": false}

# List variations in JSON format
$ dvcx variations list -p my-app -f dark-mode -o json
[
  {
    "key": "variation-on",
    "name": "Dark Mode On",
    "variables": {"enabled": true}
  },
  {
    "key": "variation-off",
    "name": "Dark Mode Off",
    "variables": {"enabled": false}
  }
]
```

---

## get

Get details of a specific variation.

### Usage

```bash
dvcx variations get <variation-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `variation-key` | The unique key of the variation |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get variation details
$ dvcx variations get variation-on -p my-app -f dark-mode
KEY:       variation-on
NAME:      Dark Mode On
VARIABLES: {"enabled": true}

# Get variation details in JSON format
$ dvcx variations get variation-on -p my-app -f dark-mode -o json
{
  "key": "variation-on",
  "name": "Dark Mode On",
  "variables": {
    "enabled": true
  }
}
```

---

## create

Create a new variation for a feature.

### Usage

```bash
dvcx variations create [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--name` | `-n` | Variation name | Yes |
| `--key` | `-k` | Variation key | Yes |
| `--variables` | `-v` | Variable values as JSON | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Create a simple variation
$ dvcx variations create -p my-app -f dark-mode \
  --key variation-dim \
  --name "Dim Mode"
KEY:       variation-dim
NAME:      Dim Mode
VARIABLES: -

# Create a variation with variables
$ dvcx variations create -p my-app -f dark-mode \
  --key variation-custom \
  --name "Custom Theme" \
  --variables '{"enabled": true, "brightness": 80}'
KEY:       variation-custom
NAME:      Custom Theme
VARIABLES: {"enabled": true, "brightness": 80}

# Create variation and output as JSON
$ dvcx variations create -p my-app -f new-checkout \
  --key v2 \
  --name "Checkout V2" \
  -v '{"version": 2}' \
  -o json
{
  "key": "v2",
  "name": "Checkout V2",
  "variables": {"version": 2}
}
```

### Notes

- Variation keys must be unique within a feature
- Variation keys can contain lowercase letters, numbers, and hyphens
- Variables should match the types defined in the feature's variable schema

---

## update

Update an existing variation.

### Usage

```bash
dvcx variations update <variation-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `variation-key` | The unique key of the variation to update |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--name` | `-n` | New variation name | No |
| `--variables` | `-v` | New variable values as JSON | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Update variation name
$ dvcx variations update variation-on -p my-app -f dark-mode \
  --name "Dark Theme Enabled"

# Update variation variables
$ dvcx variations update variation-on -p my-app -f dark-mode \
  --variables '{"enabled": true, "theme": "midnight"}'

# Update both name and variables
$ dvcx variations update variation-custom -p my-app -f dark-mode \
  --name "Custom Dark Theme" \
  --variables '{"enabled": true, "brightness": 70}'
```

### Notes

- Only the specified fields will be updated
- Variation key cannot be changed after creation

---

## delete

Delete a variation from a feature.

### Usage

```bash
dvcx variations delete <variation-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `variation-key` | The unique key of the variation to delete |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--force` | | Skip confirmation prompt | No |

### Example

```bash
# Delete a variation (with confirmation prompt)
$ dvcx variations delete variation-old -p my-app -f dark-mode
Are you sure you want to delete variation 'variation-old'? [y/N]: y
Variation 'variation-old' deleted successfully

# Delete a variation without confirmation
$ dvcx variations delete variation-test -p my-app -f dark-mode --force
Variation 'variation-test' deleted successfully
```

### Warning

- This action cannot be undone
- Variations that are actively used in targeting rules cannot be deleted
- Use `--force` flag to skip confirmation in automated scripts
