# Releasing

This document describes the release process for `dvcx`.

## Release Checklist

1. **Preview changelog** to verify release content:

   ```bash
   make changelog-preview
   ```

2. **Create and push tag**:

   ```bash
   git tag vX.Y.Z
   git push origin vX.Y.Z
   ```

3. **GitHub Actions will automatically**:
   - Generate release notes using git-cliff
   - Create GitHub Release with the notes
   - Update `CHANGELOG.md` in main branch
   - Build and publish binaries for all platforms
   - Update Homebrew formula

4. **Verify the release**:
   - Check [GitHub Releases](https://github.com/135yshr/devcycle-cli/releases)
   - Verify Homebrew formula was updated
   - Test installation: `brew upgrade dvcx`

## Changelog Generation

This project uses [git-cliff](https://git-cliff.org/) to automatically generate changelogs from gitmoji commit messages.

### Local Commands

```bash
# Generate full CHANGELOG.md
make changelog

# Preview unreleased changes
make changelog-preview
```

### Supported Gitmoji Categories

| Emoji | Category |
|-------|----------|
| ✨ | Added |
| 🐛 🩹 | Fixed |
| 📝 | Documentation |
| ⚡ | Performance |
| ♻️ 🎨 | Changed |
| ✅ 🧪 | Testing |
| 👷 🔧 📦 | Build |
| 🔒 🔐 | Security |
| ⬆️ ⬇️ 📌 | Dependencies |
| 🗑️ | Deprecated |
| 🔥 ➖ | Removed |
| 💥 | Breaking Changes |

## Version Numbering

See [Versioning Guide](docs/versioning.md) for details on version format and semantic versioning.

## Troubleshooting

### Release workflow failed

1. Check the [Actions tab](https://github.com/135yshr/devcycle-cli/actions) for error details
2. Common issues:
   - Missing `HOMEBREW_TAP_GITHUB_TOKEN` secret
   - GoReleaser configuration errors
   - git-cliff configuration errors

### CHANGELOG.md not updated

The changelog is updated in a separate job after the release. If it fails:

1. Check the "Update CHANGELOG.md" job in the workflow
2. Manually run: `make changelog && git add CHANGELOG.md && git commit -m "📝 docs: Update CHANGELOG.md"`

### Homebrew formula not updated

1. Verify `HOMEBREW_TAP_GITHUB_TOKEN` secret is set
2. Check the [homebrew-tap repository](https://github.com/135yshr/homebrew-tap) for errors
