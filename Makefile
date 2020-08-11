# PROJECT_NAME := "github-actions-demo-go"
# PKG := "github.com/brpaz/$(PROJECT_NAME)"
# PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
# GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

## GITHUB_ACTIONS is set when running as a Github Action
 
.PHONY: all lint vet test test-coverage build clean
 
all: build

dep: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@go get -u golang.org/x/lint/golint 
	@golint -set_exit_status ./...

vet: ## Run go vet
	@go vet ./...

test: ## Run unit tests
	@go test ./...

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ./... 
	@cat cover.out >> coverage.txt

build:  ## Build the binary file
	@goreleaser --snapshot --skip-publish --rm-dist
	@cp dist/FLEcli_darwin_amd64/FLEcli .
 
clean: ## Remove previous build
	@rm -f ./FLEcli
	@rm -f ./cover.out
	@rm -f ./coverage.txt
 
help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'