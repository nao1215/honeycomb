.PHONY: build test clean vet fmt chkfmt changelog tools help coverage-tree

APP         = honeycomb
VERSION     = $(shell git describe --tags --abbrev=0)
GO          = go
GO_BUILD    = $(GO) build
GO_FORMAT   = $(GO) fmt
GOFMT       = gofmt
GO_LIST     = $(GO) list
GO_TEST     = $(GO) test -v
GO_TOOL     = $(GO) tool
GO_VET      = $(GO) vet
GO_DEP      = $(GO) mod
GOOS        = ""
GOARCH      = ""
GO_PKGROOT  = ./...
GO_PACKAGES = $(shell $(GO_LIST) $(GO_PKGROOT))
GO_LDFLAGS  = -ldflags '-X github.com/nao1215/honeycomb/version.Version=${VERSION}'

build:  ## Build binary
	env GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO_BUILD) $(GO_LDFLAGS) -o $(APP) main.go

clean: ## Clean project
	-rm -rf $(APP) cover.out cover.html

test: ## Start test
	env GOOS=$(GOOS) $(GO_TEST) -cover $(GO_PKGROOT) -coverprofile=cover.out
	$(GO_TOOL) cover -html=cover.out -o cover.html

vet: ## Start go vet
	$(GO_VET) $(GO_PACKAGES)

fmt: ## Format go source code 
	$(GO_FORMAT) $(GO_PKGROOT)

coverage-tree: test ## Generate coverage tree
	go-cover-treemap -statements -coverprofile cover.out > doc/img/cover-tree.svg

changelog: ## Generate changelog. You must set GITHUB_TOKEN.
	ghch --format markdown --all --token=$(GITHUB_TOKEN) > CHANGELOG.md

tools: ## Install dependency tools 
	$(GO_INSTALL) github.com/nikolaydubina/go-cover-treemap@latest
	$(GO_INSTALL) github.com/Songmu/ghch/cmd/ghch@latest
	$(GO_INSTALL) github.com/google/wire/cmd/wire@latest

generate: ## Generate code from templates
	$(GO) generate ./...

.DEFAULT_GOAL := help
help:  
	@grep -E '^[0-9a-zA-Z_-]+[[:blank:]]*:.*?## .*$$' $(MAKEFILE_LIST) | sort \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[1;32m%-15s\033[0m %s\n", $$1, $$2}'
