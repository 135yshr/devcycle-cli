---
title: "keys"
weight: 6
---

# keys

Commands for managing SDK keys in DevCycle environments.

SDK keys are used to authenticate your application with DevCycle. Each environment has its own set of SDK keys for different platforms (client, server, mobile).

## list

List all SDK keys for an environment.

### Usage

```bash
dvcx keys list [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--environment` | `-e` | Environment key | Yes |
| `--output` | `-o` | Output format (table, json, yaml) | No |

### Example

```bash
# List SDK keys for development environment
$ dvcx keys list -p my-app -e development
ENVIRONMENT   TYPE     KEY
development   client   dvc_client_abc123...
development   server   dvc_server_def456...
development   mobile   dvc_mobile_ghi789...

# List SDK keys in JSON format
$ dvcx keys list -p my-app -e development -o json
{
  "environment": "development",
  "keys": {
    "client": "dvc_client_abc123...",
    "server": "dvc_server_def456...",
    "mobile": "dvc_mobile_ghi789..."
  }
}
```

### Key Types

| Type | Description | Use Case |
|------|-------------|----------|
| `client` | Client-side SDK key | Web browsers, frontend applications |
| `server` | Server-side SDK key | Backend services, APIs |
| `mobile` | Mobile SDK key | iOS, Android applications |

### Notes

- SDK keys are unique per environment
- Never expose server-side keys in client applications
- Client and mobile keys are safe to include in client-side code

---

## rotate

Rotate an SDK key for an environment.

### Usage

```bash
dvcx keys rotate [flags]
```

### Flags

| Flag | Short | Description | Required |
|------|-------|-------------|----------|
| `--project` | `-p` | Project key | Yes (or set in config) |
| `--environment` | `-e` | Environment key | Yes |
| `--type` | `-t` | Key type to rotate (client, server, mobile) | Yes |
| `--force` | `-f` | Skip confirmation prompt | No |

### Example

```bash
# Rotate client SDK key with confirmation
$ dvcx keys rotate -p my-app -e development --type client
Are you sure you want to rotate the client SDK key for environment 'development'?
This will invalidate the existing key. [y/N]: y
Previous Key: dvc_client_abc123...
New Key:      dvc_client_xyz987...

# Rotate server SDK key without confirmation
$ dvcx keys rotate -p my-app -e development --type server --force
Previous Key: dvc_server_def456...
New Key:      dvc_server_uvw654...

# Rotate mobile SDK key for production
$ dvcx keys rotate -p my-app -e production --type mobile --force
Previous Key: dvc_mobile_ghi789...
New Key:      dvc_mobile_rst321...
```

### Warning

**Key rotation is irreversible!**

- The old key will be immediately invalidated
- All applications using the old key will lose access to DevCycle
- Make sure to update your applications with the new key before the old key expires
- Consider performing key rotation during low-traffic periods

### Best Practices

1. **Schedule rotation**: Plan key rotations during maintenance windows
2. **Update applications first**: Have the new key ready to deploy before rotation
3. **Monitor after rotation**: Watch for authentication errors after rotation
4. **Rotate regularly**: Periodic key rotation improves security
