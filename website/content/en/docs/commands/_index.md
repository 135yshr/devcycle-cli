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

### Features

| Command | Description |
|---------|-------------|
| [features list]({{< relref "/docs/commands/features#list" >}}) | List features in a project |
| [features get]({{< relref "/docs/commands/features#get" >}}) | Get feature details |

### Variables

| Command | Description |
|---------|-------------|
| [variables list]({{< relref "/docs/commands/variables#list" >}}) | List variables in a project |
| [variables get]({{< relref "/docs/commands/variables#get" >}}) | Get variable details |

### Environments

| Command | Description |
|---------|-------------|
| [environments list]({{< relref "/docs/commands/environments#list" >}}) | List environments in a project |
| [environments get]({{< relref "/docs/commands/environments#get" >}}) | Get environment details |

### Other

| Command | Description |
|---------|-------------|
| [version]({{< relref "/docs/commands/version" >}}) | Show version information |
