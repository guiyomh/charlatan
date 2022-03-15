IGNORED_FOLDER := ".ignore"

default:
    @just --list

# Install tooling
tool: _install-lint _install-releaser
    @go install golang.org/x/tools/...@latest
    @go install github.com/joho/godotenv/...@latest
    @go install github.com/smartystreets/goconvey@latest

watch-test:
    goconvey -excludedDirs doc,example

# Build a snapshot of a project
build: _install-releaser
    goreleaser build --snapshot --rm-dist -f .goreleaser.yml

# Run all tests
test: _create-ignore
    #!/usr/bin/env bash
    go test -tags=test -v -count=1 -race -coverprofile={{IGNORED_FOLDER}}/coverage.txt -covermode=atomic $(go list ./...)
    go tool cover -func {{IGNORED_FOLDER}}/coverage.txt | grep total:

# Run units tests and generate coverage file
unit-test: _create-ignore
    #!/usr/bin/env bash
    go test -v -count=1 -race -coverprofile={{IGNORED_FOLDER}}/coverage-unit.txt -covermode=atomic $(go list ./...)
    go tool cover -func {{IGNORED_FOLDER}}/coverage-unit.txt | grep total:

spec-test: _create-ignore
    #!/usr/bin/env bash
    go test -tags=spec -v -count=1 -race -coverprofile={{IGNORED_FOLDER}}/coverage-spec.txt -covermode=atomic $(go list ./...)
    go tool cover -func {{IGNORED_FOLDER}}/coverage-spec.txt | grep total:

# Cleanup the ignored folders and binaries
clean:
    @rm -rf {{IGNORED_FOLDER}}
    @rm -rf dist/

# Build and start a mariadb server
mariadb-start:
    docker build -f example/docker/mariadb/Dockerfile -t charlatan/mariadb example/docker/mariadb
    docker run -d -p 3306:3306 --name charlatan charlatan/mariadb

# Stop and remove the mariadb server
mariadb-stop:
    docker stop charlatan
    docker rm -f charlatan

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
