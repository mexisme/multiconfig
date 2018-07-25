PACKAGE  = github.com/mexisme/multiconfig
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell \
            git describe --tags --always --dirty 2>/dev/null || \
			cat $(CURDIR)/.version 2>/dev/null)
GOPATH   = $(CURDIR)/.gopath~
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)/src/$(PACKAGE)
PKGS     = $(or $(PKG),$(shell cd $(BASE) && env GOPATH=$(GOPATH) $(GO) list ./... | grep -v "^$(PACKAGE)/vendor/"))
TESTPKGS = $(shell env GOPATH=$(GOPATH) $(GO) list -f '{{ if .TestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))

COVERAGE_PROFILE = all.coverprofile
# COVERAGE_XML = coverage.xml
COVERAGE_HTML = coverage.index.html

ifneq ($(DATE),)
LDFLAGS_DATE = -X $(PACKAGE)/version.buildDate=$(DATE)
endif
ifneq ($(VERSION),)
LDFLAGS_VERSION = -X $(PACKAGE)/version.release=$(VERSION)
endif

GO      = go
# GO_TEST = go test
GO_TEST = $(GINKGO) -r -p
GODOC   = godoc
GOFMT   = gofmt
TIMEOUT = 15
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m>>\033[0m")

.PHONY: all
all: fmt lint vendor | $(BASE) ; $(info $(M) building executable...) @ ## Build program binary
	$Q cd $(BASE) && $(GO) build \
		-tags release \
		-ldflags '$(LDFLAGS_VERSION) $(LDFLAGS_DATE)' \
		-o bin/$(PACKAGE) main.go

$(BASE): ; $(info $(M) setting GOPATH...)
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

# Tools

GOLINT = $(BIN)/gometalinter
$(GOLINT): | $(BASE) ; $(info $(M) building gometalinter...)
	$Q go get github.com/alecthomas/gometalinter
	$Q $@ --install

GODEP = $(BIN)/dep
go-dep: $(GODEP)
$(GODEP): | $(BASE) ; $(info $(M) building go-dep...)
	$Q go get -u github.com/golang/dep/cmd/dep

GINKGO = $(BIN)/ginkgo
$(GINKGO): | $(BASE) ; $(info $(M) building ginkgo...)
	$Q go get -u github.com/onsi/ginkgo/ginkgo

GOCOVMERGE = $(BIN)/gocovmerge
$(GOCOVMERGE): | $(BASE) ; $(info $(M) building gocovmerge...)
	$Q go get github.com/wadey/gocovmerge

GOCOV = $(BIN)/gocov
$(GOCOV): | $(BASE) ; $(info $(M) building gocov...)
	$Q go get github.com/axw/gocov/...

# Tests

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
INTEGRATION_TEST_TARGETS := test-integration test-integration-verbose

.PHONY: $(TEST_TARGETS) $(INTEGRATION_TEST_TARGETS) \
	check test tests run-test
## Run benchmarks
test-bench: ARGS=-run=__absolutelynothing__ -bench=.
## Run only short tests
test-short: ARGS=-short
## Run tests in verbose mode with coverage reporting
test-verbose test-integration-verbose: ARGS=-v -trace -cover
## Run tests with race detector
test-race: ARGS=-race
$(TEST_TARGETS): SKIP_ARGS=-skip=Integration
$(INTEGRATION_TEST_TARGETS): SKIP_ARGS=-focus=Integration
$(TEST_TARGETS) $(INTEGRATION_TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS) $(INTEGRATION_TEST_TARGETS): test
check test tests: fmt lint run-test
run-test: vendor | $(BASE) $(GINKGO) ; $(info $(M) running $(NAME:%=% )tests...) @ ## Run tests
	$Q cd $(BASE) && $(GO_TEST) $(ARGS) $(SKIP_ARGS) $(TESTPKGS)

# COVERAGE_PROFILE = $(COVERAGE_DIR)/all.coverprofile
# COVERAGE_XML = $(COVERAGE_DIR)/coverage.xml
# COVERAGE_HTML = $(COVERAGE_DIR)/index.html

.PHONY: cover coverage
cover coverage: COVERAGE_DIR:=$(BASE)
cover coverage: COVERAGE_FILES=$(shell find . -name '*.coverprofile')
cover coverage: vendor | $(BASE) $(GOCOVMERGE) $(GOCOV)
	$Q $(GOCOVMERGE) $(COVERAGE_FILES) > $(COVERAGE_DIR)/$(COVERAGE_PROFILE)
	$Q $(GO) tool cover -func=$(COVERAGE_DIR)/$(COVERAGE_PROFILE)
	$Q $(GO) tool cover -html=$(COVERAGE_DIR)/$(COVERAGE_PROFILE) -o $(COVERAGE_DIR)/$(COVERAGE_HTML)

# Code format

.PHONY: lint
lint: vendor | $(BASE) $(GOLINT) ; $(info $(M) running golint...) @ ## Run golint
	$Q cd $(BASE) && ret=0 && for pkg in $(PKGS); do \
		test -z "$$($(GOLINT) $$pkg | tee /dev/stderr)" || ret=1 ; \
	 done ; exit $$ret

.PHONY: fmt
fmt: ; $(info $(M) running gofmt...) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

# Dependency management

Gopkg.lock: Gopkg.toml | $(BASE) $(GODEP) ; $(info $(M) updating dependencies...)
	$Q cd $(BASE) && $(GODEP) ensure -update && $(GODEP) prune
	@touch $@
vendor: Gopkg.lock | $(BASE) $(GODEP) ; $(info $(M) retrieving dependencies...)
	$Q cd $(BASE) && $(GODEP) ensure
	@touch $@
.PHONY: go-dep-init
go-dep-init: | $(BASE) $(GODEP) ; $(info $(M) retrieving dependencies...)
	$Q cd $(BASE) && $(GODEP) init

# Misc

.PHONY: clean
clean: ; $(info $(M) cleaning...)	@ ## Cleanup everything
	@rm -rf $(GOPATH)
	@rm -rf bin
	@rm -rf test/tests.* test/coverage.*

.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
