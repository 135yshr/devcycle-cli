# Usage Guide

This guide provides detailed usage examples for all `dvcx` commands.

## Table of Contents

- [Global Flags](#global-flags)
- [Authentication](#authentication)
- [Projects](#projects)
- [Features](#features)
- [Variables](#variables)
- [Environments](#environments)
- [Targeting](#targeting)
- [Variations](#variations)
- [Audiences](#audiences)
- [Overrides](#overrides)
- [Audit Logs](#audit-logs)
- [Metrics](#metrics)
- [Webhooks](#webhooks)
- [Custom Properties](#custom-properties)

## Global Flags

All commands support these global flags:

```bash
--output, -o    Output format: table, json, yaml (default: table)
--project, -p   Project key (overrides config default)
--help, -h      Show help for command
```

## Authentication

### Login

Authenticate with DevCycle using OAuth2 credentials.

```bash
# Login using credentials from config file
dvcx auth login

# Login with explicit credentials
dvcx auth login --client-id YOUR_CLIENT_ID --client-secret YOUR_CLIENT_SECRET
```

### Logout

Remove stored authentication token.

```bash
dvcx auth logout
```

## Projects

### List Projects

```bash
# List all projects
dvcx projects list

# Output as JSON
dvcx projects list -o json
```

### Get Project

```bash
# Get project details
dvcx projects get my-project
```

### Create Project

```bash
# Create a new project
dvcx projects create --name "My Project" --key my-project --description "Project description"
```

### Update Project

```bash
# Update project name
dvcx projects update my-project --name "Updated Name"
```

## Features

### List Features

```bash
# List all features in project
dvcx features list -p my-project

# Filter by status
dvcx features list -p my-project --status active
```

### Get Feature

```bash
# Get feature details
dvcx features get my-feature -p my-project
```

### Create Feature

```bash
# Create a new feature
dvcx features create -p my-project \
  --key my-feature \
  --name "My Feature" \
  --description "Feature description" \
  --type release
```

### Update Feature

```bash
# Update feature
dvcx features update my-feature -p my-project --name "Updated Feature Name"
```

### Delete Feature

```bash
# Delete feature (with confirmation)
dvcx features delete my-feature -p my-project

# Delete without confirmation
dvcx features delete my-feature -p my-project --force
```

## Variables

### List Variables

```bash
# List all variables
dvcx variables list -p my-project
```

### Get Variable

```bash
# Get variable details
dvcx variables get my-variable -p my-project
```

### Create Variable

```bash
# Create a Boolean variable
dvcx variables create -p my-project \
  --key my-variable \
  --name "My Variable" \
  --type Boolean

# Create a String variable
dvcx variables create -p my-project \
  --key string-var \
  --name "String Variable" \
  --type String
```

### Update Variable

```bash
# Update variable
dvcx variables update my-variable -p my-project --name "Updated Name"
```

### Delete Variable

```bash
# Delete variable
dvcx variables delete my-variable -p my-project --force
```

## Environments

### List Environments

```bash
# List all environments
dvcx environments list -p my-project
```

### Get Environment

```bash
# Get environment details
dvcx environments get development -p my-project
```

## Targeting

### Get Targeting

```bash
# Get targeting configuration for a feature in an environment
dvcx targeting get -p my-project -f my-feature -e development
```

### Update Targeting

```bash
# Update targeting rules
dvcx targeting update -p my-project -f my-feature -e development \
  --status active \
  --rules '[{"audience":"all-users","variation":"on"}]'
```

### Enable/Disable Feature

```bash
# Enable feature in an environment
dvcx targeting enable -p my-project -f my-feature -e development

# Disable feature in an environment
dvcx targeting disable -p my-project -f my-feature -e development
```

## Variations

### List Variations

```bash
# List all variations for a feature
dvcx variations list -p my-project -f my-feature
```

### Get Variation

```bash
# Get variation details
dvcx variations get variation-on -p my-project -f my-feature
```

### Create Variation

```bash
# Create a new variation
dvcx variations create -p my-project -f my-feature \
  --key variation-new \
  --name "New Variation" \
  --variables '{"my-variable": true}'
```

### Update Variation

```bash
# Update variation
dvcx variations update variation-on -p my-project -f my-feature \
  --name "Updated Variation"
```

### Delete Variation

```bash
# Delete variation
dvcx variations delete variation-old -p my-project -f my-feature --force
```

## Audiences

Audiences allow you to define reusable user segments for targeting.

### List Audiences

```bash
# List all audiences
dvcx audiences list -p my-project

# Output as JSON
dvcx audiences list -p my-project -o json
```

### Get Audience

```bash
# Get audience details
dvcx audiences get beta-users -p my-project
```

### Create Audience

```bash
# Create a new audience with filters
dvcx audiences create -p my-project \
  --key beta-users \
  --name "Beta Users" \
  --description "Users in the beta program" \
  --filters '[{"type":"user","subType":"email","comparator":"contain","values":["@beta.example.com"]}]'
```

### Update Audience

```bash
# Update audience name and description
dvcx audiences update beta-users -p my-project \
  --name "Beta Testers" \
  --description "Updated description"
```

### Delete Audience

```bash
# Delete audience (with confirmation)
dvcx audiences delete beta-users -p my-project

# Delete without confirmation
dvcx audiences delete beta-users -p my-project --force
```

## Overrides

Overrides (self-targeting) allow you to set specific variations for yourself during development.

### List Overrides

```bash
# List all overrides for a feature
dvcx overrides list -p my-project -f my-feature

# List all your overrides in the project
dvcx overrides list-mine -p my-project
```

### Get Override

```bash
# Get your current override for a feature
dvcx overrides get -p my-project -f my-feature -e development
```

### Set Override

```bash
# Set an override to a specific variation
dvcx overrides set -p my-project -f my-feature -e development \
  --variation variation-on
```

### Delete Override

```bash
# Delete your override for a feature
dvcx overrides delete -p my-project -f my-feature -e development

# Delete all your overrides in the project
dvcx overrides delete-mine -p my-project
```

## Audit Logs

View audit logs for tracking changes in your project.

### List Project Audit Logs

```bash
# List all audit logs for a project
dvcx audit list -p my-project

# Output as JSON for detailed information
dvcx audit list -p my-project -o json
```

### List Feature Audit Logs

```bash
# List audit logs for a specific feature
dvcx audit feature my-feature -p my-project
```

## Metrics

Metrics allow you to measure the impact of your feature flags.

### List Metrics

```bash
# List all metrics
dvcx metrics list -p my-project
```

### Get Metric

```bash
# Get metric details
dvcx metrics get conversion-rate -p my-project
```

### Create Metric

```bash
# Create a count metric
dvcx metrics create -p my-project \
  --key page-views \
  --name "Page Views" \
  --type count \
  --event-type pageview \
  --optimize-for increase \
  --description "Tracks page view events"

# Create a conversion metric
dvcx metrics create -p my-project \
  --key signup-rate \
  --name "Signup Rate" \
  --type conversion \
  --event-type signup \
  --optimize-for increase
```

### Update Metric

```bash
# Update metric
dvcx metrics update page-views -p my-project \
  --name "Total Page Views" \
  --description "Updated description"
```

### Delete Metric

```bash
# Delete metric
dvcx metrics delete page-views -p my-project --force
```

### Get Metric Results

```bash
# Get metric results
dvcx metrics results conversion-rate -p my-project

# Filter by environment
dvcx metrics results conversion-rate -p my-project --environment production

# Filter by feature
dvcx metrics results conversion-rate -p my-project --feature my-feature

# Filter by date range
dvcx metrics results conversion-rate -p my-project \
  --start-date 2024-01-01 \
  --end-date 2024-01-31

# Combine filters
dvcx metrics results conversion-rate -p my-project \
  --environment production \
  --feature my-feature \
  --start-date 2024-01-01 \
  --end-date 2024-01-31
```

## Webhooks

Webhooks allow you to receive notifications when events occur in your project.

### List Webhooks

```bash
# List all webhooks
dvcx webhooks list -p my-project
```

### Get Webhook

```bash
# Get webhook details
dvcx webhooks get webhook-id -p my-project
```

### Create Webhook

```bash
# Create an enabled webhook
dvcx webhooks create -p my-project \
  --url "https://example.com/webhook" \
  --description "Production webhook" \
  --enabled

# Create a disabled webhook
dvcx webhooks create -p my-project \
  --url "https://example.com/webhook" \
  --description "Test webhook"
```

### Update Webhook

```bash
# Update webhook URL
dvcx webhooks update webhook-id -p my-project \
  --url "https://new-url.example.com/webhook"

# Enable webhook
dvcx webhooks update webhook-id -p my-project --enabled

# Disable webhook
dvcx webhooks update webhook-id -p my-project --disabled
```

### Delete Webhook

```bash
# Delete webhook
dvcx webhooks delete webhook-id -p my-project --force
```

## Custom Properties

Custom properties define additional user attributes for targeting.

### List Custom Properties

```bash
# List all custom properties
dvcx custom-properties list -p my-project

# Using alias
dvcx cp list -p my-project
```

### Get Custom Property

```bash
# Get custom property details
dvcx custom-properties get user-type -p my-project
```

### Create Custom Property

```bash
# Create a String property
dvcx custom-properties create -p my-project \
  --key user-type \
  --display-name "User Type" \
  --type String \
  --description "Type of user account"

# Create a Boolean property
dvcx custom-properties create -p my-project \
  --key is-premium \
  --display-name "Is Premium" \
  --type Boolean \
  --description "Whether the user is a premium subscriber"

# Create a Number property
dvcx custom-properties create -p my-project \
  --key account-age \
  --display-name "Account Age" \
  --type Number \
  --description "Days since account creation"
```

### Update Custom Property

```bash
# Update custom property
dvcx custom-properties update user-type -p my-project \
  --display-name "User Account Type" \
  --description "Updated description"
```

### Delete Custom Property

```bash
# Delete custom property
dvcx custom-properties delete user-type -p my-project --force
```

## Output Formats

All commands support multiple output formats:

```bash
# Table format (default)
dvcx features list -p my-project

# JSON format
dvcx features list -p my-project -o json

# YAML format
dvcx features list -p my-project -o yaml
```

## Configuration

### Config File

Create `.devcycle/config.yaml` in your project root:

```yaml
client_id: your-client-id
client_secret: your-client-secret
project: default-project-key
```

### Environment Variables

You can also use environment variables:

```bash
export DVCX_CLIENT_ID=your-client-id
export DVCX_CLIENT_SECRET=your-client-secret
export DVCX_PROJECT=default-project-key
```

Environment variables take precedence over config file values.
