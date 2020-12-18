.PHONY: build test help

help: ## Show this help
	printf "\033[33m%s:\033[0m\n" 'Available commands'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build you project on spaghetti-code
	go build ./cmd/apiServer

test: ## Testing project
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build
