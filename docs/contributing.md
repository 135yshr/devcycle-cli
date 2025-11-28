# Contributing Guide

Thank you for your interest in contributing to `dvcx`! This guide will help you get started.

[日本語版](ja/contributing.md)

## How to Contribute

### Reporting Issues

- Use [GitHub Issues](https://github.com/135yshr/devcycle-cli/issues) to report bugs or request features
- Check existing issues before creating a new one
- Use the provided issue templates

### Submitting Pull Requests

1. Fork the repository
2. Create a feature branch from `main`
3. Make your changes
4. Write/update tests as needed
5. Ensure all tests pass
6. Submit a pull request

## Development Workflow

### Branch Naming Convention

```
feature/<phase>-<task-name>
```

Examples:
- `feature/phase1-auth` - Authentication feature
- `feature/phase1-projects-list` - Projects list command
- `fix/token-refresh` - Bug fix for token refresh

### Commit Message Convention

We use [Gitmoji](https://gitmoji.dev/) for commit messages:

```
<gitmoji> <type>: <description>

[optional body]
```

**Common Gitmojis:**

| Emoji | Code | Description |
|-------|------|-------------|
| ✨ | `:sparkles:` | New feature |
| 🐛 | `:bug:` | Bug fix |
| 📝 | `:memo:` | Documentation |
| ♻️ | `:recycle:` | Refactoring |
| ✅ | `:white_check_mark:` | Add/update tests |
| 🔧 | `:wrench:` | Configuration |
| ⬆️ | `:arrow_up:` | Upgrade dependencies |

**Examples:**
```
✨ feat: Add projects list command
🐛 fix: Fix token expiration handling
📝 docs: Update API reference
```

### Pull Request Process

1. Fill out the PR template completely
2. Link related issues
3. Ensure CI checks pass
4. Request review from maintainers
5. Address review feedback
6. Squash commits if requested

## Code Style

### Go Code

- Follow standard Go conventions
- Run `make fmt` before committing
- Run `make lint` to check for issues
- Write tests for new functionality

### Documentation

- Write in English (primary)
- Provide Japanese translations in `docs/ja/`
- Keep README.md concise
- Put detailed information in `docs/`

## Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

## Questions?

Feel free to:
- Open an issue for questions
- Start a discussion in GitHub Discussions

## Code of Conduct

Please read and follow our [Code of Conduct](../CODE_OF_CONDUCT.md).
