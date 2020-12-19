.PHONY: build test help migrate-up migrate-down mmm

help: ## Show this help
	printf "\033[33m%s:\033[0m\n" 'Available commands'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: ## Build you project on spaghetti-code
	go build ./cmd/apiServer

test: ## Testing project
	go test -v -race -timeout 30s ./...

migrate-up: ## Migrations tables
	migrate -path migrations -database "mysql://root:myrootpassword@tcp(localhost:3306)/go-rest" up

migrate-down: ## Rollback migration tables
	migrate -path migrations -database "mysql://root:myrootpassword@tcp(localhost:3306)/go-rest" down

.DEFAULT_GOAL := build
