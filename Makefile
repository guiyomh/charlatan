SOURCE = $(wildcard *.go)
TAG ?= $(shell git describe --tags)
GOBUILD = go build -ldflags '-w'
TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

ALL = \
	$(foreach arch,x64 x32,\
	$(foreach suffix,linux osx windows,\
		build/go-faker-fixture-$(suffix)-$(arch))) \
	$(foreach arch,arm arm64,\
		build/go-faker-fixture-linux-$(arch))


help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

all: deps test build

tools: ## Install tools
	go get -u github.com/kardianos/govendor
	go get -u golang.org/x/tools/cmd/stringer
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/davecgh/go-spew/spew

build: clean test $(ALL) ## Build all binaries

test: fmtcheck
	go list $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=60s -parallel=4

clean: ## Clean Target
	rm -f $(ALL)

deps:
	govendor sync

lint:
	golint $(GOFMT_FILES)

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

vendor-status:
	@govendor status

run: clean build/go-faker-fixture-osx-x64 ## Build an Run osx
	LOG=* ./build/go-faker-fixture-osx-x64 load ./fixtures -u fixtures_user -p fixtures_pass -d fixtures

mariadb/start:
	@docker build -f docker/mariadb/Dockerfile -t gofixtures docker/mariadb
	@docker run -d -p 3306:3306 --name=gofixtures gofixtures

mariadb/stop:
	@docker rm -f gofixtures

# os is determined as thus: if variable of suffix exists, it's taken, if not, then
# suffix itself is taken
osx = darwin
build/go-faker-fixture-%-x64: $(SOURCE)
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=$(firstword $($*) $*) GOARCH=amd64 $(GOBUILD) -o $@

build/go-faker-fixture-%-x32: $(SOURCE)
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=$(firstword $($*) $*) GOARCH=386 $(GOBUILD) -o $@

build/go-faker-fixture-linux-arm: $(SOURCE)
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -o $@

build/go-faker-fixture-linux-arm64: $(SOURCE)
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -o $@
