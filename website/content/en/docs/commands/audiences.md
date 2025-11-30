---
title: "audiences"
weight: 10
---

# audiences

Commands for managing DevCycle audiences. Audiences allow you to define reusable user segments for targeting.

## list

List all audiences in a project.

### Usage

```bash
dvcx audiences list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List all audiences
$ dvcx audiences list -p my-app
KEY           NAME            DESCRIPTION           CREATED AT
beta-users    Beta Users      Beta program users    2024-01-20T10:00:00Z
premium       Premium Users   Paying customers      2024-02-15T14:30:00Z
internal      Internal Team   Company employees     2024-03-01T09:00:00Z

# List audiences in JSON format
$ dvcx audiences list -p my-app -o json
```

---

## get

Get details of a specific audience.

### Usage

```bash
dvcx audiences get <audience-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `audience-key` | The unique key of the audience |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get audience details
$ dvcx audiences get beta-users -p my-app

# Get audience details in JSON format
$ dvcx audiences get beta-users -p my-app -o json
{
  "_id": "aud-123",
  "key": "beta-users",
  "name": "Beta Users",
  "description": "Users enrolled in the beta program",
  "filters": {
    "operator": "and",
    "filters": [
      {
        "type": "user",
        "subType": "email",
        "comparator": "contain",
        "values": ["@beta.example.com"]
      }
    ]
  }
}
```

---

## create

Create a new audience.

### Usage

```bash
dvcx audiences create [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--key` | `-k` | Audience key | Yes |
| `--name` | `-n` | Audience name | Yes |
| `--description` | `-d` | Audience description | No |
| `--filters` | | Filters JSON | Yes |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Create an audience for beta users
$ dvcx audiences create -p my-app \
  --key beta-users \
  --name "Beta Users" \
  --description "Users in the beta program" \
  --filters '[{"type":"user","subType":"email","comparator":"contain","values":["@beta.example.com"]}]'

# Create an audience for premium users
$ dvcx audiences create -p my-app \
  --key premium-users \
  --name "Premium Users" \
  --filters '[{"type":"user","subType":"customData","dataKey":"plan","dataKeyType":"String","comparator":"=","values":["premium"]}]'
```

### Filter Types

| Type | SubType | Description |
|------|---------|-------------|
| `user` | `user_id` | Filter by user ID |
| `user` | `email` | Filter by email address |
| `user` | `country` | Filter by country code |
| `user` | `platform` | Filter by platform |
| `user` | `deviceModel` | Filter by device model |
| `user` | `appVersion` | Filter by app version |
| `user` | `customData` | Filter by custom data property |
| `all` | - | Match all users |

### Comparators

| Comparator | Description |
|------------|-------------|
| `=` | Equals |
| `!=` | Not equals |
| `contain` | Contains |
| `!contain` | Does not contain |
| `startWith` | Starts with |
| `!startWith` | Does not start with |
| `endWith` | Ends with |
| `!endWith` | Does not end with |
| `exist` | Exists |
| `!exist` | Does not exist |
| `>`, `>=`, `<`, `<=` | Numeric comparisons |

---

## update

Update an existing audience.

### Usage

```bash
dvcx audiences update <audience-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `audience-key` | The unique key of the audience to update |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--name` | `-n` | New audience name | No |
| `--description` | `-d` | New audience description | No |
| `--filters` | | New filters JSON | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Update audience name
$ dvcx audiences update beta-users -p my-app --name "Beta Testers"

# Update audience description
$ dvcx audiences update beta-users -p my-app --description "Updated description"

# Update audience filters
$ dvcx audiences update beta-users -p my-app \
  --filters '[{"type":"user","subType":"email","comparator":"contain","values":["@beta.example.com","@test.example.com"]}]'
```

---

## delete

Delete an audience from a project.

### Usage

```bash
dvcx audiences delete <audience-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `audience-key` | The unique key of the audience to delete |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--force` | | Skip confirmation prompt | No |

### Example

```bash
# Delete an audience (with confirmation)
$ dvcx audiences delete beta-users -p my-app
Are you sure you want to delete audience 'beta-users'? [y/N]: y
Audience 'beta-users' deleted successfully

# Delete without confirmation
$ dvcx audiences delete beta-users -p my-app --force
```

### Notes

- Deleting an audience may affect features that reference it in targeting rules
- This action cannot be undone
