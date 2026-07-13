.PHONY: setup dev lint format format-check test build

setup:
	mise trust --yes
	mise install
	lefthook install

dev:
	@go run .

lint:
	@golangci-lint run ./...

format:
	@golangci-lint fmt ./...

format-check:
	@golangci-lint fmt ./... --diff

test:
	@go test ./... -cover

build:
	@rm -rf ./bin
	@mkdir -p ./bin
	@go build -o ./bin/main main.go

%:
	@:
