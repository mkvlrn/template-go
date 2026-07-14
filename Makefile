.DEFAULT_GOAL := build

.PHONY: setup dev lint format format-check test build

setup:
	mise trust --yes
	mise install
	lefthook install

dev:
	@go run .

lint:
	@golangci-lint run --config=./.config/golangci.toml ./...

format:
	@gofumpt -w .

format-check:
	@gofumpt -d .

test:
	@go test ./... -cover

build:
	@rm -rf ./bin
	@mkdir -p ./bin
	@go build -o ./bin/main main.go

%:
	@:
