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

---

## create

Create a new project in your DevCycle organization.

### Usage

```bash
dvcx projects create [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--name` | `-n` | Project name | Yes |
| `--key` | `-k` | Project key | Yes |
| `--description` | `-d` | Project description | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Create a new project
$ dvcx projects create -n "My New App" -k my-new-app
KEY:         my-new-app
NAME:        My New App
CREATED:     2024-06-20T10:00:00Z

# Create a project with description
$ dvcx projects create -n "Production App" -k production-app -d "Main production application"

# Create project and output as JSON
$ dvcx projects create -n "Staging App" -k staging-app -o json
{
  "key": "staging-app",
  "name": "Staging App",
  "createdAt": "2024-06-20T10:00:00Z"
}
```

### Notes

- Project keys must be unique within your organization
- Project keys can contain lowercase letters, numbers, and hyphens
- Once created, project keys cannot be changed

---

## update

Update an existing project.

### Usage

```bash
dvcx projects update <project-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `project-key` | The unique key of the project to update |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--name` | `-n` | New project name | No |
| `--description` | `-d` | New project description | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Update project name
$ dvcx projects update my-app -n "My Updated Application"

# Update project description
$ dvcx projects update my-app -d "Updated description for my application"

# Update both name and description
$ dvcx projects update my-app -n "New Name" -d "New description"
```

### Notes

- Only the specified fields will be updated
- Project key cannot be changed after creation
