# Security

This document describes the security measures implemented in `dvcx`.

## Authentication

### OAuth2 Token Management

- Tokens are obtained via OAuth2 client credentials flow from `https://auth.devcycle.com/oauth/token`
- Tokens are stored locally in `.devcycle/token` with restricted file permissions (0600)
- The `.devcycle/` directory is created with restricted permissions (0700)
- Tokens are automatically refreshed when expired
- All API communication uses HTTPS

### Credential Storage

- Client credentials (`client_id` and `client_secret`) are stored in `.devcycle/config.yaml`
- The `.devcycle/` directory is excluded from version control via `.gitignore`
- Tokens are stored in plaintext but with restricted file permissions

## Input Validation

### URL Path Escaping

All user-provided path parameters are escaped using `url.PathEscape()` to prevent:

- Path traversal attacks
- URL injection
- Special character manipulation

This applies to all API endpoints including:

- Project keys
- Feature keys
- Environment keys
- Variable keys
- Webhook IDs
- Metric keys
- Audience keys
- Custom property keys

### Query Parameter Escaping

Query parameters are escaped using `url.QueryEscape()` to prevent injection attacks.

### File Size Limits

To prevent denial-of-service attacks via memory exhaustion:

- JSON input files are limited to 10MB maximum
- Stdin input is limited using `io.LimitReader`
- File size is checked before reading

## Network Security

### HTTPS Only

- All API communication uses HTTPS
- Base URL: `https://api.devcycle.com/v1`
- Authentication URL: `https://auth.devcycle.com/oauth/token`

### Request Timeouts

- All API requests have a 30-second timeout
- Prevents hanging connections

## Best Practices for Users

### Protect Your Credentials

1. **Never commit credentials** - The `.devcycle/` directory is in `.gitignore`
2. **Use environment variables** - For CI/CD, use `DVCX_CLIENT_ID` and `DVCX_CLIENT_SECRET`
3. **Rotate credentials** - Periodically rotate your DevCycle client credentials

### Secure Your Environment

1. **File permissions** - Ensure `.devcycle/` has restricted permissions
2. **Shared systems** - Be cautious when using `dvcx` on shared systems
3. **CI/CD secrets** - Use your CI/CD platform's secret management

## Reporting Security Issues

If you discover a security vulnerability, please report it by creating a private security advisory on GitHub or contacting the maintainers directly.

**Do not** create public issues for security vulnerabilities.
