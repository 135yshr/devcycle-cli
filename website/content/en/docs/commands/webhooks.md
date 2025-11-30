---
title: "webhooks"
weight: 22
---

# webhooks

Commands for managing DevCycle webhooks. Webhooks allow you to receive notifications when events occur in your project.

## list

List all webhooks in a project.

### Usage

```bash
dvcx webhooks list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List all webhooks
$ dvcx webhooks list -p my-app
ID            URL                                  ENABLED    CREATED AT
wh-123        https://example.com/webhook1         Yes        2024-06-20T10:00:00Z
wh-456        https://example.com/webhook2         No         2024-06-21T14:30:00Z

# List webhooks in JSON format
$ dvcx webhooks list -p my-app -o json
```

---

## get

Get details of a specific webhook.

### Usage

```bash
dvcx webhooks get <webhook-id> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `webhook-id` | The unique ID of the webhook |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Get webhook details
$ dvcx webhooks get wh-123 -p my-app

# Get webhook details in JSON format
$ dvcx webhooks get wh-123 -p my-app -o json
{
  "_id": "wh-123",
  "url": "https://example.com/webhook",
  "description": "Production notifications",
  "isEnabled": true,
  "createdAt": "2024-06-20T10:00:00Z",
  "updatedAt": "2024-06-20T10:00:00Z"
}
```

---

## create

Create a new webhook.

### Usage

```bash
dvcx webhooks create [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--url` | | Webhook URL | Yes |
| `--description` | | Webhook description | No |
| `--enabled` | | Enable the webhook (default: true) | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Create an enabled webhook
$ dvcx webhooks create -p my-app \
  --url "https://example.com/webhook" \
  --description "Production notifications" \
  --enabled

# Create a disabled webhook (for testing)
$ dvcx webhooks create -p my-app \
  --url "https://staging.example.com/webhook" \
  --description "Staging webhook"
```

### Webhook Payload

When an event occurs, DevCycle sends a POST request to your webhook URL with a JSON payload containing:

- Event type
- Project information
- Feature/variable details (if applicable)
- User who triggered the change
- Timestamp

---

## update

Update an existing webhook.

### Usage

```bash
dvcx webhooks update <webhook-id> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `webhook-id` | The unique ID of the webhook to update |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--url` | | New webhook URL | No |
| `--description` | | New webhook description | No |
| `--enabled` | | Enable the webhook | No |
| `--disabled` | | Disable the webhook | No |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# Update webhook URL
$ dvcx webhooks update wh-123 -p my-app \
  --url "https://new-url.example.com/webhook"

# Enable a webhook
$ dvcx webhooks update wh-123 -p my-app --enabled

# Disable a webhook
$ dvcx webhooks update wh-123 -p my-app --disabled

# Update description
$ dvcx webhooks update wh-123 -p my-app \
  --description "Updated description"
```

### Notes

- You cannot use both `--enabled` and `--disabled` flags together
- Only specified fields will be updated

---

## delete

Delete a webhook from a project.

### Usage

```bash
dvcx webhooks delete <webhook-id> [flags]
```

### Arguments

| Argument | Description |
|----------|-------------|
| `webhook-id` | The unique ID of the webhook to delete |

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--force` | | Skip confirmation prompt | No |

### Example

```bash
# Delete a webhook (with confirmation)
$ dvcx webhooks delete wh-123 -p my-app
Are you sure you want to delete webhook 'wh-123'? [y/N]: y
Webhook 'wh-123' deleted successfully

# Delete without confirmation
$ dvcx webhooks delete wh-123 -p my-app --force
```

### Notes

- Deleting a webhook will stop all notifications to that URL
- This action cannot be undone
