---
title: "audit"
weight: 20
---

# audit

Commands for viewing audit logs. Audit logs track all changes made to your DevCycle project.

## list

List all audit logs for a project.

### Usage

```bash
dvcx audit list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List project audit logs
$ dvcx audit list -p my-app
TYPE              USER                 CREATED AT
feature.created   john@example.com     2024-06-20T10:00:00Z
feature.updated   jane@example.com     2024-06-20T11:30:00Z
variable.created  john@example.com     2024-06-20T12:00:00Z

# List audit logs in JSON format
$ dvcx audit list -p my-app -o json
[
  {
    "_id": "audit-123",
    "type": "feature.created",
    "user": {
      "name": "John Doe",
      "email": "john@example.com"
    },
    "changes": [...],
    "createdAt": "2024-06-20T10:00:00Z"
  }
]
```

### Audit Log Types

| Type | Description |
|------|-------------|
| `feature.created` | A feature was created |
| `feature.updated` | A feature was updated |
| `feature.deleted` | A feature was deleted |
| `variable.created` | A variable was created |
| `variable.updated` | A variable was updated |
| `variable.deleted` | A variable was deleted |
| `targeting.updated` | Targeting rules were updated |
| `variation.created` | A variation was created |
| `variation.updated` | A variation was updated |
| `variation.deleted` | A variation was deleted |

---

## feature

List audit logs for a specific feature.

### Usage

```bash
dvcx audit feature <feature-key> [flags]
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
# List audit logs for a specific feature
$ dvcx audit feature dark-mode -p my-app
TYPE              USER                 CREATED AT
feature.created   john@example.com     2024-06-20T10:00:00Z
targeting.updated jane@example.com     2024-06-20T11:30:00Z
variation.created john@example.com     2024-06-20T12:00:00Z

# List feature audit logs in JSON format
$ dvcx audit feature dark-mode -p my-app -o json
```

### Notes

- Audit logs are read-only and cannot be modified
- Logs include details about what was changed, who made the change, and when
- Use audit logs to track changes and troubleshoot issues
