---
title: "auth"
weight: 1
---

# auth

Authentication commands for DevCycle API.

## login

Authenticate with DevCycle using OAuth2 credentials.

### Usage

```bash
dvcx auth login
```

### Description

This command prompts you to enter your DevCycle API credentials:

- **Client ID**: Your DevCycle API client ID
- **Client Secret**: Your DevCycle API client secret

After successful authentication, the access token is stored in `.devcycle/token.json` in the current directory.

### Example

```bash
$ dvcx auth login
Enter Client ID: your-client-id
Enter Client Secret: ********
Successfully authenticated!
```

### Notes

- Credentials are obtained from the DevCycle dashboard under **Settings** → **API Credentials**
- The token is stored locally and used for subsequent API calls
- Token expiration is handled automatically; you may need to re-authenticate periodically

---

## logout

Remove stored authentication credentials.

### Usage

```bash
dvcx auth logout
```

### Description

This command removes the stored access token from `.devcycle/token.json`.

### Example

```bash
$ dvcx auth logout
Successfully logged out!
```

### Notes

- After logout, you will need to run `dvcx auth login` again before using other commands
- This only removes the local token; it does not revoke the token on the server
