.PHONY: setup dev lint format format-check test build

setup:
	mise trust --yes
	mise install
	lefthook install

dev:
	@go run ./cmd/$(filter-out $@,$(MAKECMDGOALS))

lint:
	@golangci-lint run --config=./.config/golangci.toml --default=standard ./...

format:
	@golangci-lint fmt --config=./.config/golangci.toml ./...

format-check:
	@golangci-lint fmt --config=./.config/golangci.toml ./... --diff

test:
	@go test ./...

build:
	@mkdir -p ./bin
	for dir in ./cmd/*/ ; do \
		go build -o ./bin/$$(basename $$dir) $$dir; \
	done

%:
	@:
