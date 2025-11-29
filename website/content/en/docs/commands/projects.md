---
title: "projects"
weight: 2
---

# projects

Commands for managing DevCycle projects.

## list

List all projects in your DevCycle organization.

### Usage

```bash
dvcx projects list [flags]
```

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--output` | `-o` | Output format (table, json, yaml) |

### Example

```bash
# List projects in table format (default)
$ dvcx projects list
KEY             NAME                    CREATED
my-app          My Application          2024-01-15
staging-app     Staging Application     2024-02-20
production      Production App          2024-03-10

# List projects in JSON format
$ dvcx projects list -o json
[
  {
    "key": "my-app",
    "name": "My Application",
    "createdAt": "2024-01-15T10:00:00Z"
  },
  ...
]

# List projects in YAML format
$ dvcx projects list -o yaml
- key: my-app
  name: My Application
  createdAt: 2024-01-15T10:00:00Z
...
```

---

## get

Get details of a specific project.

### Usage

```bash
dvcx projects get <project-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `project-key` | The unique key of the project |

### Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--output` | `-o` | Output format (table, json, yaml) |

### Example

```bash
# Get project details in table format
$ dvcx projects get my-app
KEY:         my-app
NAME:        My Application
DESCRIPTION: Main application project
CREATED:     2024-01-15T10:00:00Z

# Get project details in JSON format
$ dvcx projects get my-app -o json
{
  "key": "my-app",
  "name": "My Application",
  "description": "Main application project",
  "createdAt": "2024-01-15T10:00:00Z",
  "updatedAt": "2024-06-01T15:30:00Z"
}
```

### Notes

- The project key is case-sensitive
- Use `projects list` to see all available project keys
