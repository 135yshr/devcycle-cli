# Roadmap

This document outlines the implementation phases for `dvcx`.

## Phase Overview

| Phase | Focus | Status |
|-------|-------|--------|
| Phase 1 | Foundation + Basic Operations (MVP) | In Progress |
| Phase 2 | Write Operations | Planned |
| Phase 3 | Targeting & Variations | Planned |
| Phase 4 | Audiences & Overrides | Planned |
| Phase 5 | Operations & Monitoring | Planned |
| Phase 6 | Environment Management | Planned |

---

## Phase 1: Foundation + Basic Operations (MVP)

**Goal:** Build CLI foundation and implement most frequently used read operations.

### Infrastructure
- [x] Go module initialization
- [x] CLI framework (cobra) setup
- [x] Configuration management (viper)
- [ ] HTTP client implementation
- [ ] Output formatters (table/JSON/YAML)

### Authentication
- [ ] `auth login` - OAuth2 token acquisition
- [ ] `auth logout` - Token removal

### Read Commands
- [ ] `projects list` - List all projects
- [ ] `projects get` - Get project details
- [ ] `features list` - List all features
- [ ] `features get` - Get feature details
- [ ] `variables list` - List all variables
- [ ] `variables get` - Get variable details
- [ ] `environments list` - List all environments
- [ ] `environments get` - Get environment details

---

## Phase 2: Write Operations

**Goal:** Implement create, update, and delete operations for features and variables.

### Feature Management
- [ ] `features create` - Create a new feature
- [ ] `features update` - Update a feature
- [ ] `features delete` - Delete a feature

### Variable Management
- [ ] `variables create` - Create a new variable
- [ ] `variables update` - Update a variable
- [ ] `variables delete` - Delete a variable

### Project Management
- [ ] `projects create` - Create a new project
- [ ] `projects update` - Update a project

---

## Phase 3: Targeting & Variations

**Goal:** Implement feature flag configuration features.

### Targeting
- [ ] `targeting get` - Get targeting configuration
- [ ] `targeting update` - Update targeting rules

### Variations
- [ ] `variations list` - List all variations
- [ ] `variations create` - Create a new variation
- [ ] `variations update` - Update a variation
- [ ] `variations delete` - Delete a variation

---

## Phase 4: Audiences & Overrides

**Goal:** Implement advanced targeting features.

### Audiences
- [ ] `audiences list` - List all audiences
- [ ] `audiences get` - Get audience details
- [ ] `audiences create` - Create a new audience
- [ ] `audiences update` - Update an audience
- [ ] `audiences delete` - Delete an audience

### Overrides (Self-Targeting)
- [ ] `overrides list` - List all overrides
- [ ] `overrides create` - Create a new override
- [ ] `overrides delete` - Delete an override

---

## Phase 5: Operations & Monitoring

**Goal:** Implement operational features like audit logs, metrics, and webhooks.

### Audit Logs
- [ ] `audit list` - List project audit logs
- [ ] `audit feature` - List feature audit logs

### Metrics
- [ ] `metrics project` - Get project metrics
- [ ] `metrics feature` - Get feature metrics

### Webhooks
- [ ] `webhooks list` - List all webhooks
- [ ] `webhooks create` - Create a new webhook
- [ ] `webhooks update` - Update a webhook
- [ ] `webhooks delete` - Delete a webhook

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
