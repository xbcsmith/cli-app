PACKAGE  = github.com/xbcsmith/cli-app/examples/go-nodeinfo
BINARY    = ./bin/nodeinfo
COMMIT  ?= $(shell git rev-parse --short=16 HEAD)

gitversion := $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo 0.1.0-0)
VERSION ?= $(gitversion)

TOOLS    = $(CURDIR)/tools
GO      = go
GOFMT   = gofmt
GOBUILD = CGO_ENABLED=0 go build -v

all: fmt lint test $(BINARY)

SOURCES = $(shell find -name vendor -prune -o -name \*.go -print)

$(BINARY): $(SOURCES); $(info $(M) building executable…) @ ## Build program binary
	$Q $(GOBUILD) $(TAGS) -o $@ .

GOLINT = $(TOOLS)/golint
$(GOLINT): ; $(info $(M) building golint…)
	$Q go build -o $@ golang.org/x/lint/golint


.PHONY: lint
lint: $(GOLINT) ; $(info $(M) running golint…) @ ## Run golint change ret=1 to make lint required
	$Q ret=0 && for pkg in $(PKGS); do \
		test -z "$$($(GOLINT) $$pkg | tee /dev/stderr)" || ret=0 ; \
	 done ; exit $$ret

.PHONY: fmt
fmt: ; $(info $(M) running gofmt…) @ ## Run gofmt on all source files
	@ret=0 && for d in $$($(GO) list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		$(GOFMT) -l -w $$d/*.go || ret=$$? ; \
	 done ; exit $$ret

.PHONY: test
test: ; $(info $(M) running tests…) @
	$Q go test -v

.PHONY: clean
clean: ; $(info $(M) cleaning…)	@ ## Cleanup everything
	@rm -rf bin tools vendor


.PHONY: help
help:
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: version
version:
	@echo $(VERSION)
