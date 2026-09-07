# template-go

[![ci](https://img.shields.io/github/actions/workflow/status/mkvlrn/template-go/checks.yml?branch=main&style=flat&logo=github&label=ci)](https://github.com/mkvlrn/template-go/actions/workflows/checks.yml?query=branch%3Amain)
[![template](https://img.shields.io/badge/template-use_this_template-2ea44f?style=flat&logo=github)](https://github.com/mkvlrn/template-go/generate)
[![mise](https://mise-versions.jdx.dev/badge.svg)](https://mise.jdx.dev)
[![license](https://img.shields.io/github/license/mkvlrn/template-go?style=flat)](https://github.com/mkvlrn/template-go/blob/main/LICENSE)

A sane, opinionated template for Go projects.

> [!NOTE]
> This template includes an Arch Linux Dev Container based on [mise-devcontainers](https://github.com/mkvlrn/mise-devcontainers), providing a consistent development environment with [mise](https://mise.jdx.dev) preconfigured.
>
> `mise` manages the Go toolchain and development tools inside the container. Project tasks intentionally use `make`, keeping the usual interface expected in Go projects.

Uses, among other tools:

- [golangci-lint](https://golangci-lint.run) for linting
- [gofumpt](https://github.com/mvdan/gofumpt) for formatting
- [Lefthook](https://github.com/evilmartians/lefthook) for Git hooks
- [Cocogitto](https://github.com/cocogitto/cocogitto) for commit message linting
- [GNU Make](https://www.gnu.org/software/make/) as the task interface

## requirements and dependencies

To use the included Dev Container you need:

- Docker or a compatible container runtime
- a Dev Container-compatible editor or the [Dev Container CLI](https://github.com/devcontainers/cli)
- an SSH agent exposed through `SSH_AUTH_SOCK` with at least one key loaded

The SSH agent is forwarded into the container for Git authentication and commit signing. Private keys remain on the host.

The Go toolchain and development tools are managed by `mise` inside the container. `make` is used to run project tasks.

If you prefer not to use the Dev Container, install [mise](https://mise.jdx.dev) locally and run:

```sh
mise install
```

> [!NOTE]
> Git hooks keep the tooling managed by mise synchronized after checkouts and merges.

## running

Project tasks are exposed through the `Makefile` rather than mise tasks.

### `make run`

Runs the project.

### `make test`

Runs the test suite.

### `make lint`

Runs the configured linters.

### `make fmt`

Formats the project.

### `make build`

Builds the project.

## ci

CI is provided by GitHub Actions through [`.github/workflows/checks.yml`](https://github.com/mkvlrn/template-go/blob/main/.github/workflows/checks.yml).

It runs the project's formatting, linting, testing, and build checks.

## license

[MIT](https://github.com/mkvlrn/template-go/blob/main/LICENSE)
