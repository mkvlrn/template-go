# template-go

A sane, opinionated template for Go projects with a consistent dev environment powered by mise.

> [!NOTE]
> This template requires [mise](https://mise.jdx.dev) as the only global dependency. All instructions assume mise is installed.

mise handles tool versioning and task running, replacing both version managers and Makefiles with a single, cross-platform tool.

## uses

- **[golangci-lint](https://golangci-lint.run)** for linting and formatting
- **[lefthook](https://github.com/evilmartians/lefthook)** for git hooks
- **[mise](https://mise.jdx.dev)** for tool management and task running

## setup

```bash
mise run setup
```

Installs all tools (Go, golangci-lint, lefthook) and configures git hooks.

## tasks

`mise run dev <name>`

Runs a program from `./cmd/<name>`.

`mise run build`

Builds all programs from `./cmd` into `./bin`.

`mise run test`

Runs all tests.

`mise run lint`

Runs golangci-lint with the `standard` linter set.

`mise run format`

Formats code using gofumpt via golangci-lint.

`mise run format-check`

Checks formatting without modifying files. Exits non-zero if anything needs formatting.

## ci

GitHub Actions runs on pushes and pull requests to `main`, executing lint and tests.

## vscode

You might want to install the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go). The `.editorconfig` file handles basic formatting rules.
