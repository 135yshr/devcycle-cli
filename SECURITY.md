# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |

## Reporting a Vulnerability

We take security vulnerabilities seriously. If you discover a security issue, please report it responsibly.

### How to Report

**Please DO NOT report security vulnerabilities through public GitHub issues.**

Instead, please report them via one of the following methods:

1. **GitHub Security Advisories** (Preferred)
   - Go to the [Security Advisories](https://github.com/135yshr/devcycle-cli/security/advisories) page
   - Click "Report a vulnerability"
   - Provide detailed information about the vulnerability

2. **Email**
   - Send an email with details to the repository owner
   - Include "SECURITY" in the subject line

### What to Include

Please include the following information in your report:

- Type of vulnerability (e.g., command injection, credential exposure, etc.)
- Step-by-step instructions to reproduce the issue
- Affected versions
- Potential impact of the vulnerability
- Any suggested fixes (optional)

### What to Expect

- **Acknowledgment**: We will acknowledge receipt of your report within 48 hours
- **Updates**: We will provide updates on our progress within 7 days
- **Resolution**: We aim to resolve critical vulnerabilities within 30 days

### Safe Harbor

We consider security research conducted in accordance with this policy to be:

- Authorized in accordance with any applicable anti-hacking laws
- Exempt from restrictions in our Terms of Service that would interfere with conducting security research

We will not pursue legal action against researchers who:

- Act in good faith
- Avoid privacy violations, data destruction, or service interruption
- Report findings confidentially and allow reasonable time for remediation

## Security Best Practices for Users

When using `dvcx`, please follow these security best practices:

### Credential Management

- **Never commit credentials** to version control
- Store `client_id` and `client_secret` in `.devcycle/config.yaml` (already in `.gitignore`)
- Use environment variables (`DVCX_CLIENT_ID`, `DVCX_CLIENT_SECRET`) for CI/CD environments
- Rotate API credentials periodically

### Configuration File Security

```bash
# Ensure config file has appropriate permissions
chmod 600 .devcycle/config.yaml
```

### Network Security

- Always use HTTPS endpoints (enforced by the tool)
- Be cautious when using the tool on shared or public networks

## Acknowledgments

We appreciate the security research community's efforts in helping keep `dvcx` and its users safe.
