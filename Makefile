.PHONY: build install clean test lint fmt help

BINARY_NAME=dvcx
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS=-ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

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

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^## //p' $(MAKEFILE_LIST) | column -t -s ':' | sed 's/^/  /'
