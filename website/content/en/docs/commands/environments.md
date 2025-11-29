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
