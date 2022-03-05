IGNORED_FOLDER := ".ignore"
COVERAGE_FILE := "coverage.txt"
GOFILES := `go list ./...`

default:
    @just --list

# Install tooling
tool: _install-lint _install-releaser
    @go install golang.org/x/tools/...@latest
    @go install github.com/joho/godotenv/...@latest

# Build a snapshot of a project
build: _install-releaser
    goreleaser build --snapshot --rm-dist -f .goreleaser.yml

# Run units tests and generate coverage file
unit-test: _create-ignore
    go test -v -count=1 -race -coverprofile={{IGNORED_FOLDER}}/{{COVERAGE_FILE}} -covermode=atomic {{GOFILES}}
    go tool cover -func {{IGNORED_FOLDER}}/{{COVERAGE_FILE}} | grep total:

# Cleanup the ignored folders and binaries
clean:
    @rm -rf {{IGNORED_FOLDER}}
    @rm -rf dist/

#Run lint step with golangci-lint
lint: _install-lint
    golangci-lint run

_install-lint:
    @go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.41.1

_install-releaser:
    @go install github.com/goreleaser/goreleaser/...@latest

_create-ignore:
    @if [ ! -d {{IGNORED_FOLDER}} ]; then \
        mkdir -p {{IGNORED_FOLDER}}; \
    fi
