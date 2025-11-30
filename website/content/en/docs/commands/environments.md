---
title: "environments"
weight: 5
---

# environments

Commands for managing DevCycle environments.

Environments represent different deployment stages (development, staging, production) where your features can have different configurations.

**Aliases**: `envs`, `env`

## list

List all environments in a project.

### Usage

```bash
dvcx environments list [flags]
# or
dvcx envs list [flags]
dvcx env list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List environments in a project
$ dvcx environments list -p my-app
KEY           NAME          TYPE
development   Development   development
staging       Staging       staging
production    Production    production

# List environments in JSON format
$ dvcx envs list -p my-app -o json
[
  {
    "key": "development",
    "name": "Development",
    "type": "development"
  },
  {
    "key": "staging",
    "name": "Staging",
    "type": "staging"
  },
  {
    "key": "production",
    "name": "Production",
    "type": "production"
  }
]
```

### Environment Types

| Type | Description |
|------|-------------|
| `development` | Development environment for testing |
| `staging` | Staging environment for pre-production testing |
| `production` | Production environment for live users |

---

## get

Get details of a specific environment.

### Usage

```bash
dvcx environments get <environment-key> [flags]
# or
dvcx envs get <environment-key> [flags]
dvcx env get <environment-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `environment-key` | The unique key of the environment |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get environment details
$ dvcx environments get production -p my-app
KEY:         production
NAME:        Production
TYPE:        production
DESCRIPTION: Live production environment
CREATED:     2024-01-15T10:00:00Z

# Get environment details in JSON format
$ dvcx env get production -p my-app -o json
{
  "key": "production",
  "name": "Production",
  "type": "production",
  "description": "Live production environment",
  "createdAt": "2024-01-15T10:00:00Z",
  "updatedAt": "2024-01-15T10:00:00Z"
}
```

### Notes

- Each project can have multiple environments
- Environment configurations are independent - a feature can be enabled in development but disabled in production
- SDK keys are unique per environment

---

## create

Create a new environment in a project.

### Usage

```bash
dvcx environments create [flags]
# or
dvcx envs create [flags]
dvcx env create [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--name` | `-n` | Environment name | Yes |
| `--key` | `-k` | Environment key | Yes |
| `--type` | `-t` | Environment type (development, staging, production) | No (default: development) |
| `--color` | | Environment color (hex format) | No |
| `--description` | `-d` | Environment description | No |

### Example

```bash
# Create a staging environment
$ dvcx environments create -p my-app \
  --key staging \
  --name "Staging" \
  --type staging \
  --color "#ffff00" \
  --description "Pre-production testing environment"
KEY:         staging
NAME:        Staging
TYPE:        staging
DESCRIPTION: Pre-production testing environment
CREATED:     2024-01-15T10:00:00Z

# Create a minimal environment
$ dvcx envs create -p my-app --key qa --name "QA"
KEY:         qa
NAME:        QA
TYPE:        development
CREATED:     2024-01-15T10:00:00Z
```

---

## update

Update an existing environment.

### Usage

```bash
dvcx environments update <environment-key> [flags]
# or
dvcx envs update <environment-key> [flags]
dvcx env update <environment-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `environment-key` | The unique key of the environment to update |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--name` | `-n` | New environment name | No |
| `--color` | | New environment color (hex format) | No |
| `--description` | `-d` | New environment description | No |

### Example

```bash
# Update environment name and description
$ dvcx environments update staging -p my-app \
  --name "Staging Environment" \
  --description "Updated staging environment"
KEY:         staging
NAME:        Staging Environment
TYPE:        staging
DESCRIPTION: Updated staging environment
UPDATED:     2024-01-16T10:00:00Z

# Update only the color
$ dvcx envs update staging -p my-app --color "#00ffff"
KEY:         staging
NAME:        Staging Environment
TYPE:        staging
COLOR:       #00ffff
UPDATED:     2024-01-16T10:00:00Z
```

---

## delete

Delete an environment from a project.

### Usage

```bash
dvcx environments delete <environment-key> [flags]
# or
dvcx envs delete <environment-key> [flags]
dvcx env delete <environment-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `environment-key` | The unique key of the environment to delete |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--force` | `-f` | Skip confirmation prompt | No |

### Example

```bash
# Delete environment with confirmation
$ dvcx environments delete staging -p my-app
Are you sure you want to delete environment 'staging'? [y/N]: y
Environment deleted successfully

# Delete environment without confirmation
$ dvcx envs delete qa -p my-app --force
Environment deleted successfully
```

### Warning

Deleting an environment is irreversible and will remove all feature configurations for that environment.
