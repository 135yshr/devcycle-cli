---
title: "custom-properties"
weight: 23
---

# custom-properties

Commands for managing custom properties. Custom properties define additional user attributes that can be used for targeting.

**Alias:** `cp`

## list

List all custom properties in a project.

### Usage

```bash
dvcx custom-properties list [flags]
# or
dvcx cp list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List all custom properties
$ dvcx custom-properties list -p my-app
KEY           DISPLAY NAME     TYPE      DESCRIPTION              CREATED AT
user-type     User Type        String    Type of user account     2024-06-20T10:00:00Z
is-premium    Is Premium       Boolean   Premium subscriber       2024-06-21T14:30:00Z
account-age   Account Age      Number    Days since creation      2024-06-22T09:00:00Z

# Using alias
$ dvcx cp list -p my-app

# List in JSON format
$ dvcx cp list -p my-app -o json
```

---

## get

Get details of a specific custom property.

### Usage

```bash
dvcx custom-properties get <property-key> [flags]
# or
dvcx cp get <property-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `property-key` | The unique key of the custom property |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get custom property details
$ dvcx cp get user-type -p my-app

# Get in JSON format
$ dvcx cp get user-type -p my-app -o json
{
  "_id": "cp-123",
  "key": "user-type",
  "propertyKey": "user-type",
  "displayName": "User Type",
  "type": "String",
  "description": "Type of user account"
}
```

---

## create

Create a new custom property.

### Usage

```bash
dvcx custom-properties create [flags]
# or
dvcx cp create [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--key` | `-k` | Property key (must match SDK property name) | Yes |
| `--display-name` | | Display name | Yes |
| `--type` | `-t` | Property type (Boolean, Number, String) | Yes |
| `--description` | | Property description | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Create a String property
$ dvcx cp create -p my-app \
  --key user-type \
  --display-name "User Type" \
  --type String \
  --description "Type of user account"

# Create a Boolean property
$ dvcx cp create -p my-app \
  --key is-premium \
  --display-name "Is Premium" \
  --type Boolean \
  --description "Whether the user is a premium subscriber"

# Create a Number property
$ dvcx cp create -p my-app \
  --key account-age \
  --display-name "Account Age" \
  --type Number \
  --description "Days since account creation"
```

### Property Types

| Type | Description | Example Values |
|------|-------------|----------------|
| `Boolean` | True/false values | `true`, `false` |
| `Number` | Numeric values | `42`, `3.14`, `-10` |
| `String` | Text values | `"premium"`, `"basic"` |

### Notes

- The property key must match the property name used in your SDK implementation
- Property keys can contain lowercase letters, numbers, hyphens, and underscores
- Property type cannot be changed after creation

---

## update

Update an existing custom property.

### Usage

```bash
dvcx custom-properties update <property-key> [flags]
# or
dvcx cp update <property-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `property-key` | The unique key of the custom property to update |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--display-name` | | New display name | No |
| `--description` | | New description | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Update display name
$ dvcx cp update user-type -p my-app \
  --display-name "User Account Type"

# Update description
$ dvcx cp update user-type -p my-app \
  --description "Updated description"

# Update both
$ dvcx cp update user-type -p my-app \
  --display-name "User Account Type" \
  --description "The type of user account (basic, premium, enterprise)"
```

### Notes

- Property key and type cannot be changed after creation
- Only display name and description can be updated

---

## delete

Delete a custom property from a project.

### Usage

```bash
dvcx custom-properties delete <property-key> [flags]
# or
dvcx cp delete <property-key> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `property-key` | The unique key of the custom property to delete |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--force` | | Skip confirmation prompt | No |

### Example

```bash
# Delete a custom property (with confirmation)
$ dvcx cp delete user-type -p my-app
Are you sure you want to delete custom property 'user-type'? [y/N]: y
Custom property 'user-type' deleted successfully

# Delete without confirmation
$ dvcx cp delete user-type -p my-app --force
```

### Notes

- Deleting a custom property may affect targeting rules that use it
- This action cannot be undone
- Consider updating targeting rules before deleting a property

## Using Custom Properties in Targeting

Once created, custom properties can be used in audience filters and targeting rules:

```bash
# Create an audience using custom property
$ dvcx audiences create -p my-app \
  --key premium-users \
  --name "Premium Users" \
  --filters '[{"type":"user","subType":"customData","dataKey":"user-type","dataKeyType":"String","comparator":"=","values":["premium"]}]'
```

### SDK Usage

Custom properties must be passed from your application code:

```javascript
// JavaScript SDK example
const user = {
  user_id: "user-123",
  customData: {
    "user-type": "premium",
    "is-premium": true,
    "account-age": 365
  }
};
```

Make sure the property keys in your SDK match the keys defined in DevCycle.
