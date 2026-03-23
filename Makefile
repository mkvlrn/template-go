.PHONY: setup dev lint format format-check test build

setup:
	mise trust --yes
	mise install
	lefthook install

dev:
	go run ./cmd/$(filter-out $@,$(MAKECMDGOALS))

lint:
	golangci-lint run --default=standard ./...

format:
	golangci-lint fmt ./...

format-check:
	golangci-lint fmt ./... --diff

test:
	go test ./...

build:
	@mkdir -p ./bin
	for dir in ./cmd/*/ ; do \
		go build -o ./bin/$$(basename $$dir) $$dir; \
	done

%:
	@:
