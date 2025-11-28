.PHONY: build install clean test lint fmt help

BINARY_NAME=dvcx

# Version: Extract from git tag (v1.0.0 -> 1.0.0), fallback to "dev"
GIT_TAG=$(shell git describe --tags --abbrev=0 2>/dev/null | sed 's/^v//')
VERSION=$(if $(GIT_TAG),$(GIT_TAG),dev)
# Commit: Short git hash
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
# Date: Build timestamp in ISO 8601 format
DATE=$(shell date -u '+%Y-%m-%dT%H:%M:%SZ')

# ldflags to embed version info into binary
LDFLAGS=-ldflags "\
	-s -w \
	-X github.com/135yshr/devcycle-cli/cmd.Version=$(VERSION) \
	-X github.com/135yshr/devcycle-cli/cmd.Commit=$(COMMIT) \
	-X github.com/135yshr/devcycle-cli/cmd.Date=$(DATE)"

## build: Build the binary
build:
	go build $(LDFLAGS) -o bin/$(BINARY_NAME) .

## install: Install the binary to $GOPATH/bin
install:
	go install $(LDFLAGS) .

## clean: Remove build artifacts
clean:
	rm -rf bin/
	go clean

## test: Run tests
test:
	go test -v ./...

## test-coverage: Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

## lint: Run linter
lint:
	golangci-lint run

## fmt: Format code
fmt:
	go fmt ./...
	goimports -w .

## tidy: Tidy go modules
tidy:
	go mod tidy

## run: Run the CLI
run:
	go run . $(ARGS)

## version: Show version info
version:
	@echo "Version: $(VERSION)"
	@echo "Commit:  $(COMMIT)"
	@echo "Date:    $(DATE)"

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^## //p' $(MAKEFILE_LIST) | column -t -s ':' | sed 's/^/  /'
