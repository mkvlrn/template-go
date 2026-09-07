.DEFAULT_GOAL := build

.PHONY: setup dev lint format format-check test build sync-branch

setup:
	mise trust --yes
	mise install
	lefthook install

dev:
	@go run . $(ARGS)

lint:
	@golangci-lint run ./...

format:
	@gofumpt -w .

test:
	@go test ./... -cover

build:
	@rm -rf ./bin
	@mkdir -p ./bin
	@go build -o ./bin/main main.go


sync-branch:
	@if command -v mise >/dev/null 2>&1; then mise prune -y; fi
	@go mod tidy
	@lefthook install

%:
	@:
