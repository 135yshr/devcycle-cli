---
title: "targeting"
weight: 6
---

# targeting

Commands for managing DevCycle feature targeting configurations.

Targeting allows you to control which users see specific variations of your feature flags based on rules, percentages, and user attributes.

## get

Get the targeting configuration for a feature.

### Usage

```bash
dvcx targeting get [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get targeting configuration for a feature
$ dvcx targeting get -p my-app -f dark-mode
ENVIRONMENT   STATUS    TARGETS
development   active    2 rule(s)
staging       active    1 rule(s)
production    inactive  0 rule(s)

# Get targeting configuration in JSON format
$ dvcx targeting get -p my-app -f dark-mode -o json
{
  "development": {
    "status": "active",
    "targets": [...]
  },
  "staging": {
    "status": "active",
    "targets": [...]
  },
  "production": {
    "status": "inactive",
    "targets": []
  }
}
```

---

## update

Update the targeting configuration for a feature using a JSON file.

### Usage

```bash
dvcx targeting update [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--from-file` | `-F` | JSON input file for configuration, use '-' for stdin | Yes |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Update targeting from a JSON file
$ dvcx targeting update -p my-app -f dark-mode -F targeting-config.json

# Update targeting from stdin
$ cat targeting-config.json | dvcx targeting update -p my-app -f dark-mode -F -

# Example targeting-config.json structure:
{
  "development": {
    "status": "active",
    "targets": [
      {
        "name": "All Users",
        "distribution": [
          {"_variation": "variation-on", "percentage": 1.0}
        ],
        "audience": {
          "name": "All Users",
          "filters": {
            "filters": [{"type": "all"}],
            "operator": "and"
          }
        }
      }
    ]
  }
}
```

### Notes

- The JSON file must contain a map of environment keys to configuration objects
- Maximum file size is 10MB
- Use stdin (`-F -`) for piping configurations from other commands

---

## enable

Enable a feature for a specific environment.

### Usage

```bash
dvcx targeting enable [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--environment` | `-e` | Environment key | Yes |

### Example

```bash
# Enable a feature for the development environment
$ dvcx targeting enable -p my-app -f dark-mode -e development
Feature 'dark-mode' enabled for environment 'development'

# Enable a feature for production
$ dvcx targeting enable -p my-app -f new-checkout -e production
Feature 'new-checkout' enabled for environment 'production'
```

### Notes

- Enabling a feature activates its targeting rules for the specified environment
- Users will start receiving variations based on the configured targeting rules

---

## disable

Disable a feature for a specific environment.

### Usage

```bash
dvcx targeting disable [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--feature` | `-f` | Feature key | Yes |
| `--environment` | `-e` | Environment key | Yes |

### Example

```bash
# Disable a feature for the production environment
$ dvcx targeting disable -p my-app -f dark-mode -e production
Feature 'dark-mode' disabled for environment 'production'

# Disable a feature for staging
$ dvcx targeting disable -p my-app -f experimental-feature -e staging
Feature 'experimental-feature' disabled for environment 'staging'
```

### Notes

- Disabling a feature stops serving it to users in the specified environment
- Users will receive the default/off variation when the feature is disabled
- Targeting rules are preserved and will be applied again when the feature is re-enabled
