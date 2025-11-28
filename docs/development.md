# Development Guide

This guide covers setting up your development environment and understanding the project architecture.

## Prerequisites

- Go 1.24 or later
- Make
- Git

## Setup

1. **Clone the repository**

   ```bash
   git clone https://github.com/135yshr/devcycle-cli.git
   cd devcycle-cli
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Build the project**

   ```bash
   make build
   ```

4. **Run tests**

   ```bash
   make test
   ```

## Build Commands

| Command | Description |
|---------|-------------|
| `make build` | Build binary to `bin/dvcx` |
| `make install` | Install to `$GOPATH/bin` |
| `make test` | Run all tests |
| `make test-coverage` | Run tests with coverage report |
| `make lint` | Run golangci-lint |
| `make fmt` | Format code |
| `make tidy` | Tidy go modules |
| `make clean` | Remove build artifacts |
| `make run ARGS="..."` | Run CLI with arguments |

## Project Structure

```text
devcycle-cli/
├── cmd/                    # Cobra command definitions
│   ├── root.go            # Root command and global flags
│   ├── auth.go            # Authentication commands
│   ├── projects.go        # Project commands
│   ├── features.go        # Feature commands
│   └── ...
├── internal/              # Internal packages
│   ├── api/               # API client
│   │   ├── client.go      # HTTP client
│   │   ├── auth.go        # Authentication
│   │   └── ...
│   ├── config/            # Configuration
│   │   └── config.go      # Viper configuration
│   └── output/            # Output formatting
│       └── formatter.go   # Table/JSON/YAML formatters
├── docs/                  # Documentation
│   ├── api-reference.md   # API endpoint reference
│   ├── roadmap.md         # Implementation roadmap
│   ├── contributing.md    # Contribution guide
│   ├── development.md     # This file
│   └── ja/                # Japanese translations
├── main.go                # Entry point
├── go.mod                 # Go module definition
├── Makefile               # Build tasks
└── README.md              # Project overview
```

## Architecture

### Command Layer (`cmd/`)

Each resource type has its own file containing Cobra commands:

```go
// cmd/projects.go
var projectsCmd = &cobra.Command{
    Use:   "projects",
    Short: "Manage DevCycle projects",
}

var projectsListCmd = &cobra.Command{
    Use:   "list",
    Short: "List all projects",
    RunE:  runProjectsList,
}
```

### API Layer (`internal/api/`)

HTTP client for DevCycle Management API:

```go
// internal/api/client.go
type Client struct {
    baseURL    string
    httpClient *http.Client
    token      string
}

func (c *Client) Get(path string) (*Response, error)
func (c *Client) Post(path string, body interface{}) (*Response, error)
```

### Configuration (`internal/config/`)

Uses Viper for configuration management:

- Config file: `.devcycle/config.yaml`
- Token file: `.devcycle/token.json`
- Environment variables: `DVCX_*`

### Output Formatting (`internal/output/`)

Supports multiple output formats:

- Table (default) - Human-readable
- JSON - Machine-readable
- YAML - Configuration-friendly

## Adding a New Command

1. Create command file in `cmd/`:

   ```go
   // cmd/newresource.go
   var newresourceCmd = &cobra.Command{
       Use:   "newresource",
       Short: "Manage new resources",
   }

   func init() {
       rootCmd.AddCommand(newresourceCmd)
   }
   ```

2. Add API methods in `internal/api/`:

   ```go
   // internal/api/newresource.go
   func (c *Client) ListNewResources() ([]NewResource, error)
   func (c *Client) GetNewResource(id string) (*NewResource, error)
   ```

3. Write tests

4. Update documentation

## Testing

### Unit Tests

```bash
go test ./...
```

### Integration Tests

Integration tests require DevCycle API credentials:

```bash
DVCX_CLIENT_ID=xxx DVCX_CLIENT_SECRET=xxx go test ./... -tags=integration
```

## Debugging

Run with debug output:

```bash
DVCX_DEBUG=true ./bin/dvcx projects list
```

Or add to config:

```yaml
debug: true
```
