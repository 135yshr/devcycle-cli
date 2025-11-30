---
title: "Commands"
weight: 2
bookCollapseSection: true
---

# Commands

Complete reference for all dvcx commands.

## Global Flags

These flags are available for all commands:

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--output` | `-o` | Output format (table, json, yaml) | table |
| `--config` | | Path to config file | .devcycle/config.yaml |
| `--help` | `-h` | Help for any command | |

## Command Categories

### Authentication

| Command | Description |
|---------|-------------|
| [auth login]({{< relref "/docs/commands/auth#login" >}}) | Authenticate with DevCycle |
| [auth logout]({{< relref "/docs/commands/auth#logout" >}}) | Remove stored credentials |

### Projects

| Command | Description |
|---------|-------------|
| [projects list]({{< relref "/docs/commands/projects#list" >}}) | List all projects |
| [projects get]({{< relref "/docs/commands/projects#get" >}}) | Get project details |
| [projects create]({{< relref "/docs/commands/projects#create" >}}) | Create a new project |
| [projects update]({{< relref "/docs/commands/projects#update" >}}) | Update a project |

### Features

| Command | Description |
|---------|-------------|
| [features list]({{< relref "/docs/commands/features#list" >}}) | List features in a project |
| [features get]({{< relref "/docs/commands/features#get" >}}) | Get feature details |
| [features create]({{< relref "/docs/commands/features#create" >}}) | Create a new feature |
| [features update]({{< relref "/docs/commands/features#update" >}}) | Update a feature |
| [features delete]({{< relref "/docs/commands/features#delete" >}}) | Delete a feature |

### Variables

| Command | Description |
|---------|-------------|
| [variables list]({{< relref "/docs/commands/variables#list" >}}) | List variables in a project |
| [variables get]({{< relref "/docs/commands/variables#get" >}}) | Get variable details |
| [variables create]({{< relref "/docs/commands/variables#create" >}}) | Create a new variable |
| [variables update]({{< relref "/docs/commands/variables#update" >}}) | Update a variable |
| [variables delete]({{< relref "/docs/commands/variables#delete" >}}) | Delete a variable |

### Environments

| Command | Description |
|---------|-------------|
| [environments list]({{< relref "/docs/commands/environments#list" >}}) | List environments in a project |
| [environments get]({{< relref "/docs/commands/environments#get" >}}) | Get environment details |
| [environments create]({{< relref "/docs/commands/environments#create" >}}) | Create a new environment |
| [environments update]({{< relref "/docs/commands/environments#update" >}}) | Update an environment |
| [environments delete]({{< relref "/docs/commands/environments#delete" >}}) | Delete an environment |

### Targeting

| Command | Description |
|---------|-------------|
| [targeting get]({{< relref "/docs/commands/targeting#get" >}}) | Get targeting configuration |
| [targeting update]({{< relref "/docs/commands/targeting#update" >}}) | Update targeting rules |
| [targeting enable]({{< relref "/docs/commands/targeting#enable" >}}) | Enable a feature for an environment |
| [targeting disable]({{< relref "/docs/commands/targeting#disable" >}}) | Disable a feature for an environment |

### Variations

| Command | Description |
|---------|-------------|
| [variations list]({{< relref "/docs/commands/variations#list" >}}) | List variations for a feature |
| [variations get]({{< relref "/docs/commands/variations#get" >}}) | Get variation details |
| [variations create]({{< relref "/docs/commands/variations#create" >}}) | Create a new variation |
| [variations update]({{< relref "/docs/commands/variations#update" >}}) | Update a variation |
| [variations delete]({{< relref "/docs/commands/variations#delete" >}}) | Delete a variation |

### Audiences

| Command | Description |
|---------|-------------|
| [audiences list]({{< relref "/docs/commands/audiences#list" >}}) | List all audiences |
| [audiences get]({{< relref "/docs/commands/audiences#get" >}}) | Get audience details |
| [audiences create]({{< relref "/docs/commands/audiences#create" >}}) | Create a new audience |
| [audiences update]({{< relref "/docs/commands/audiences#update" >}}) | Update an audience |
| [audiences delete]({{< relref "/docs/commands/audiences#delete" >}}) | Delete an audience |

### Overrides

| Command | Description |
|---------|-------------|
| [overrides list]({{< relref "/docs/commands/overrides#list" >}}) | List overrides for a feature |
| [overrides get]({{< relref "/docs/commands/overrides#get" >}}) | Get your override for a feature |
| [overrides set]({{< relref "/docs/commands/overrides#set" >}}) | Set an override for yourself |
| [overrides delete]({{< relref "/docs/commands/overrides#delete" >}}) | Delete your override |
| [overrides list-mine]({{< relref "/docs/commands/overrides#list-mine" >}}) | List all your overrides |
| [overrides delete-mine]({{< relref "/docs/commands/overrides#delete-mine" >}}) | Delete all your overrides |

### Audit Logs

| Command | Description |
|---------|-------------|
| [audit list]({{< relref "/docs/commands/audit#list" >}}) | List project audit logs |
| [audit feature]({{< relref "/docs/commands/audit#feature" >}}) | List feature audit logs |

### Metrics

| Command | Description |
|---------|-------------|
| [metrics list]({{< relref "/docs/commands/metrics#list" >}}) | List all metrics |
| [metrics get]({{< relref "/docs/commands/metrics#get" >}}) | Get metric details |
| [metrics create]({{< relref "/docs/commands/metrics#create" >}}) | Create a new metric |
| [metrics update]({{< relref "/docs/commands/metrics#update" >}}) | Update a metric |
| [metrics delete]({{< relref "/docs/commands/metrics#delete" >}}) | Delete a metric |
| [metrics results]({{< relref "/docs/commands/metrics#results" >}}) | Get metric results |

### Webhooks

| Command | Description |
|---------|-------------|
| [webhooks list]({{< relref "/docs/commands/webhooks#list" >}}) | List all webhooks |
| [webhooks get]({{< relref "/docs/commands/webhooks#get" >}}) | Get webhook details |
| [webhooks create]({{< relref "/docs/commands/webhooks#create" >}}) | Create a new webhook |
| [webhooks update]({{< relref "/docs/commands/webhooks#update" >}}) | Update a webhook |
| [webhooks delete]({{< relref "/docs/commands/webhooks#delete" >}}) | Delete a webhook |

### Custom Properties

| Command | Description |
|---------|-------------|
| [custom-properties list]({{< relref "/docs/commands/custom-properties#list" >}}) | List all custom properties |
| [custom-properties get]({{< relref "/docs/commands/custom-properties#get" >}}) | Get custom property details |
| [custom-properties create]({{< relref "/docs/commands/custom-properties#create" >}}) | Create a new custom property |
| [custom-properties update]({{< relref "/docs/commands/custom-properties#update" >}}) | Update a custom property |
| [custom-properties delete]({{< relref "/docs/commands/custom-properties#delete" >}}) | Delete a custom property |

### SDK Keys

| Command | Description |
|---------|-------------|
| [keys list]({{< relref "/docs/commands/keys#list" >}}) | List SDK keys for an environment |
| [keys rotate]({{< relref "/docs/commands/keys#rotate" >}}) | Rotate an SDK key |

### Other

| Command | Description |
|---------|-------------|
| [version]({{< relref "/docs/commands/version" >}}) | Show version information |
