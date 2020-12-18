.PHONY: build
build:
	go build ./cmd/apiServer

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build