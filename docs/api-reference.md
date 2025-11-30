# DevCycle Management API Reference

This document lists all DevCycle Management API endpoints that `dvcx` supports or plans to support.

## Base URL

```text
https://api.devcycle.com/v1
```

## Authentication

The Management API uses OAuth2 for authentication.

**Token Endpoint:** `https://auth.devcycle.com/oauth/token`

To obtain credentials:

1. Go to [DevCycle Dashboard](https://app.devcycle.com/)
2. Navigate to Settings > API Credentials
3. Create a new API Key with appropriate permissions

## Endpoints

### Authentication Endpoint

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/oauth/token` | POST | Obtain OAuth2 access token |

### Projects

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects` | GET | List all projects |
| `/v1/projects` | POST | Create a new project |
| `/v1/projects/{project}` | GET | Get project details |
| `/v1/projects/{project}` | PATCH | Update a project |
| `/v1/projects/{project}` | DELETE | Delete a project |

### Environments

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/environments` | GET | List all environments |
| `/v1/projects/{project}/environments` | POST | Create a new environment |
| `/v1/projects/{project}/environments/{environment}` | GET | Get environment details |
| `/v1/projects/{project}/environments/{environment}` | PATCH | Update an environment |
| `/v1/projects/{project}/environments/{environment}` | DELETE | Delete an environment |
| `/v1/projects/{project}/environments/{environment}/keys` | POST | Generate/rotate SDK keys |

### Features

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/features` | GET | List all features |
| `/v1/projects/{project}/features` | POST | Create a new feature |
| `/v1/projects/{project}/features/{feature}` | GET | Get feature details |
| `/v1/projects/{project}/features/{feature}` | PATCH | Update a feature |
| `/v1/projects/{project}/features/{feature}` | DELETE | Delete a feature |
| `/v1/projects/{project}/features/{feature}/configurations` | GET | Get feature configurations |
| `/v1/projects/{project}/features/{feature}/configurations` | PATCH | Update feature configurations (targeting) |

### Variables

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/variables` | GET | List all variables |
| `/v1/projects/{project}/variables` | POST | Create a new variable |
| `/v1/projects/{project}/variables/{variable}` | GET | Get variable details |
| `/v1/projects/{project}/variables/{variable}` | PATCH | Update a variable |
| `/v1/projects/{project}/variables/{variable}` | DELETE | Delete a variable |

### Variations

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/features/{feature}/variations` | GET | List all variations |
| `/v1/projects/{project}/features/{feature}/variations` | POST | Create a new variation |
| `/v1/projects/{project}/features/{feature}/variations/{variation}` | GET | Get variation details |
| `/v1/projects/{project}/features/{feature}/variations/{variation}` | PATCH | Update a variation |
| `/v1/projects/{project}/features/{feature}/variations/{variation}` | DELETE | Delete a variation |

### Audiences

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/audiences` | GET | List all audiences |
| `/v1/projects/{project}/audiences` | POST | Create a new audience |
| `/v1/projects/{project}/audiences/{audience}` | GET | Get audience details |
| `/v1/projects/{project}/audiences/{audience}` | PATCH | Update an audience |
| `/v1/projects/{project}/audiences/{audience}` | DELETE | Delete an audience |

### Overrides (Self-Targeting)

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/features/{feature}/overrides` | GET | List all overrides for a feature |
| `/v1/projects/{project}/features/{feature}/overrides/current` | GET | Get current user's override for a feature |
| `/v1/projects/{project}/features/{feature}/overrides/current` | PUT | Create/update override for current user |
| `/v1/projects/{project}/features/{feature}/overrides/current` | DELETE | Delete override for current user |
| `/v1/projects/{project}/overrides/current` | GET | List all overrides for current user in project |
| `/v1/projects/{project}/overrides/current` | DELETE | Delete all overrides for current user in project |

### Custom Properties

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/customProperties` | GET | List all custom properties |
| `/v1/projects/{project}/customProperties` | POST | Create a new custom property |
| `/v1/projects/{project}/customProperties/{key}` | GET | Get custom property details |
| `/v1/projects/{project}/customProperties/{key}` | PATCH | Update a custom property |
| `/v1/projects/{project}/customProperties/{key}` | DELETE | Delete a custom property |

### User Profiles

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/userProfile/current` | GET | Get current user profile |
| `/v1/projects/{project}/userProfile/current` | PUT | Create/update user profile |

### Webhooks

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/webhooks` | GET | List all webhooks |
| `/v1/projects/{project}/webhooks` | POST | Create a new webhook |
| `/v1/projects/{project}/webhooks/{webhook}` | GET | Get webhook details |
| `/v1/projects/{project}/webhooks/{webhook}` | PATCH | Update a webhook |
| `/v1/projects/{project}/webhooks/{webhook}` | DELETE | Delete a webhook |

### Audit Logs

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/audit` | GET | Get project audit logs |
| `/v1/projects/{project}/features/{feature}/audit` | GET | Get feature audit logs |

### Metrics

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/metrics` | GET | List all metrics |
| `/v1/projects/{project}/metrics` | POST | Create a new metric |
| `/v1/projects/{project}/metrics/{key}` | GET | Get metric details |
| `/v1/projects/{project}/metrics/{key}` | PATCH | Update a metric |
| `/v1/projects/{project}/metrics/{key}` | DELETE | Delete a metric |
| `/v1/projects/{project}/metrics/{key}/results` | GET | Fetch metric results |
| `/v1/projects/{project}/test-metric-results` | GET | Test metric results |

### Metric Associations

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/v1/projects/{project}/metric-associations` | GET | List metric associations |
| `/v1/projects/{project}/metric-associations` | POST | Associate metric with feature |
| `/v1/projects/{project}/metric-associations` | DELETE | Delete metric association |

## Official Documentation

For detailed request/response schemas and more information, see the [official DevCycle Management API documentation](https://docs.devcycle.com/management-api/).
