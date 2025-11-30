# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Feat: Implement Phase 6 environment management (#16)
- Feat: Implement Phase 6 environment management

### Documentation

- Docs: Add comprehensive GoDoc comments to pkg/api (#17)
- Docs: Add comprehensive GoDoc comments to pkg/api

## [0.2.0] - 2025-11-30

### Added

- Feat: Implement Phase 5 operations and monitoring (#14)
- Feat: Implement Phase 5 operations and monitoring
- Feat: Implement Phase 4 audiences and overrides (#13)
- Feat: Implement Phase 4 audiences and overrides
- Feat: Implement Phase 3 targeting and variations (#11)
- Feat: Implement Phase 3 targeting and variations

### Build

- Ci: Sync docs deployment with binary release (#15)
- Ci: Sync docs deployment with binary release

### Documentation

- Docs: Add Phase 4 & 5 command pages to website
- Docs: Add usage guide for Phase 4 and Phase 5 commands
- Docs: Fix API reference consistency with DevCycle official API (#12)
- Docs: Fix API reference consistency with DevCycle official API
- Docs: Add Homebrew installation instructions (#10)
- Docs: Add Homebrew installation instructions

### Fixed

- Fix: Use highlight shortcode in tabs for proper code rendering (#9)
- Fix: Use highlight shortcode in tabs for proper code rendering

## [0.1.1] - 2025-11-29

### Added

- Feat: Implement Phase 2 write operations (#7)
- Feat: Add stdin support for feature creation from JSON
- Feat: Add v2 API support for feature creation with JSON file input
- Feat: Implement Phase 2 write operations
- Feat: Implement Phase 1 API client and CLI commands
- Feat: Add version command with build info
- Feat: Initialize project structure with cobra/viper

### Build

- Ci: Improve CI workflow
- Ci: Add GoReleaser workflow for Homebrew tap
- Ci: Add GitHub Actions workflow for CI

### Changed

- Refactor: Move api package to pkg/ for public access

### Dependencies

- ⬆️ deps: Upgrade Hugo to 0.152.2 (latest) for hugo-book compatibility (#5)
- ⬆️ deps: Upgrade Hugo to 0.152.2 (latest) for hugo-book compatibility
- ⬆️ deps: Update hugo-book theme to v13 (#3)
- ⬆️ deps: Update hugo-book theme to v13
- ⬆️ chore: Upgrade Go to 1.25.4 and use go.mod for CI

### Documentation

- Docs: Add Phase 2 write operations to website documentation
- Docs: Add Hugo documentation site with bilingual support
- Docs: Add OSS community documents
- Docs: Change install method from pip/npm to brew
- Docs: Add development prerequisites to README
- Docs: Fix markdown lint errors
- Docs: Add comprehensive project documentation

### Fixed

- Fix: Correct Hugo module configuration for hugo-book theme (#4)
- Fix: Correct Hugo module configuration for hugo-book theme

### Other

- Style: Improve code snippet readability in dark mode (#6)
- Style: Improve code snippet readability in dark mode
- Phase 1: Implement API client and CLI commands (#2)
- ️ refactor: Rename config functions to follow Go naming conventions
- ️ refactor: Rename API methods to follow Go naming conventions
- Phase 1: Initialize project structure and documentation (#1)
- Fix duplicate CI runs on feature branches
- Add path filters to CI workflow for efficiency
- ️ perf: Optimize CI workflow to reduce jobs
- Chore: Add pre-commit configuration
- Initial commit

### Testing

- Test: Add unit tests for api, output, and config packages

<!-- Links -->
[Unreleased]: https://github.com/135yshr/devcycle-cli/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/135yshr/devcycle-cli/compare/v0.1.1...v0.2.0
[0.1.1]: https://github.com/135yshr/devcycle-cli/releases/tag/v0.1.1
