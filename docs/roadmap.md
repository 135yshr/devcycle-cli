# Roadmap

This document outlines the implementation phases for `dvcx`.

## Phase Overview

| Phase | Focus | Status |
|-------|-------|--------|
| Phase 1 | Foundation + Basic Operations (MVP) | Completed |
| Phase 2 | Write Operations | Completed |
| Phase 3 | Targeting & Variations | Completed |
| Phase 4 | Audiences & Overrides | Completed |
| Phase 5 | Operations & Monitoring | Planned |
| Phase 6 | Environment Management | Planned |

---

## Phase 1: Foundation + Basic Operations (MVP)

**Goal:** Build CLI foundation and implement most frequently used read operations.

### Infrastructure

- [x] Go module initialization
- [x] CLI framework (cobra) setup
- [x] Configuration management (viper)
- [x] HTTP client implementation
- [x] Output formatters (table/JSON/YAML)

### Authentication

- [x] `auth login` - OAuth2 token acquisition
- [x] `auth logout` - Token removal

### Read Commands

- [x] `projects list` - List all projects
- [x] `projects get` - Get project details
- [x] `features list` - List all features
- [x] `features get` - Get feature details
- [x] `variables list` - List all variables
- [x] `variables get` - Get variable details
- [x] `environments list` - List all environments
- [x] `environments get` - Get environment details

---

## Phase 2: Write Operations

**Goal:** Implement create, update, and delete operations for features and variables.

### Feature Management

- [x] `features create` - Create a new feature
- [x] `features update` - Update a feature
- [x] `features delete` - Delete a feature

### Variable Management

- [x] `variables create` - Create a new variable
- [x] `variables update` - Update a variable
- [x] `variables delete` - Delete a variable

### Project Management

- [x] `projects create` - Create a new project
- [x] `projects update` - Update a project

---

## Phase 3: Targeting & Variations

**Goal:** Implement feature flag configuration features.

### Targeting

- [x] `targeting get` - Get targeting configuration
- [x] `targeting update` - Update targeting rules
- [x] `targeting enable` - Enable a feature for an environment
- [x] `targeting disable` - Disable a feature for an environment

### Variations

- [x] `variations list` - List all variations
- [x] `variations get` - Get variation details
- [x] `variations create` - Create a new variation
- [x] `variations update` - Update a variation
- [x] `variations delete` - Delete a variation

---

## Phase 4: Audiences & Overrides

**Goal:** Implement advanced targeting features.

### Audiences

- [x] `audiences list` - List all audiences
- [x] `audiences get` - Get audience details
- [x] `audiences create` - Create a new audience
- [x] `audiences update` - Update an audience
- [x] `audiences delete` - Delete an audience

### Overrides (Self-Targeting)

- [x] `overrides list` - List all overrides for a feature (requires --feature)
- [x] `overrides get` - Get current user's override for a feature
- [x] `overrides set` - Create/update override for current user
- [x] `overrides delete` - Delete override for current user
- [x] `overrides list-mine` - List all my overrides in project
- [x] `overrides delete-mine` - Delete all my overrides in project

---

## Phase 5: Operations & Monitoring

**Goal:** Implement operational features like audit logs, metrics, webhooks, and custom properties.

### Audit Logs

- [ ] `audit list` - List project audit logs
- [ ] `audit feature` - List feature audit logs

### Metrics

- [ ] `metrics list` - List all metrics
- [ ] `metrics get` - Get metric details
- [ ] `metrics create` - Create a new metric
- [ ] `metrics update` - Update a metric
- [ ] `metrics delete` - Delete a metric
- [ ] `metrics results` - Get metric results

### Webhooks

- [ ] `webhooks list` - List all webhooks
- [ ] `webhooks get` - Get webhook details
- [ ] `webhooks create` - Create a new webhook
- [ ] `webhooks update` - Update a webhook
- [ ] `webhooks delete` - Delete a webhook

### Custom Properties

- [ ] `custom-properties list` - List all custom properties
- [ ] `custom-properties get` - Get custom property details
- [ ] `custom-properties create` - Create a new custom property
- [ ] `custom-properties update` - Update a custom property
- [ ] `custom-properties delete` - Delete a custom property

---

## Phase 6: Environment Management

**Goal:** Implement environment creation, deletion, and SDK key management.

### Environments

- [ ] `environments create` - Create a new environment
- [ ] `environments update` - Update an environment
- [ ] `environments delete` - Delete an environment

### SDK Keys

- [ ] `keys list` - List SDK keys
- [ ] `keys rotate` - Rotate SDK keys

---

## Contributing

Want to help implement these features? See our [Contributing Guide](contributing.md) for how to get started.
