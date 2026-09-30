# template-go

[![ci](https://img.shields.io/github/actions/workflow/status/mkvlrn/template-go/checks.yml?branch=main&style=flat&logo=github&label=ci)](https://github.com/mkvlrn/template-go/actions/workflows/checks.yml?query=branch%3Amain)
[![template](https://img.shields.io/badge/template-use_this_template-2ea44f?style=flat&logo=github)](https://github.com/mkvlrn/template-go/generate)
[![mise](https://mise-versions.jdx.dev/badge.svg)](https://mise.jdx.dev)
[![license](https://img.shields.io/github/license/mkvlrn/template-go?style=flat)](https://github.com/mkvlrn/template-go/blob/main/LICENSE)

A sane, opinionated template for Go projects.

> [!TIP]
> Using [mise](https://mise.jdx.dev) locally is the path of least friction: it manages the Go toolchain, development tools, and tasks without requiring a container.
>
> This template also includes an optional Arch Linux Dev Container based on [mise-devcontainers](https://github.com/mkvlrn/mise-devcontainers). It is suggested if you do not normally use `mise` or want a consistent, preconfigured development environment.

Uses, among other tools:

- [golangci-lint](https://golangci-lint.run) for linting and running the [gofumpt](https://github.com/mvdan/gofumpt) formatter
- [Lefthook](https://github.com/evilmartians/lefthook) for Git hooks
- [Cocogitto](https://github.com/cocogitto/cocogitto) for commit message linting
- [mise](https://mise.jdx.dev) as the task interface

## requirements and dependencies

Using `mise` locally is recommended and provides the least-friction setup. If you do not normally use `mise`, the included Dev Container is an optional way to get a consistent development environment with everything preconfigured.

To use the Dev Container you need:

- Docker or a compatible container runtime
- a Dev Container-compatible editor or the [Dev Container CLI](https://github.com/devcontainers/cli)
- an SSH agent exposed through `SSH_AUTH_SOCK` with at least one key loaded

The SSH agent is forwarded into the container for Git authentication and commit signing. Private keys remain on the host.

The Go toolchain and development tools are managed by `mise`. `mise` tasks are used to run project operations.

If you do not use the Dev Container, install [mise](https://mise.jdx.dev) locally and run:

```sh
mise install
```

> [!NOTE]
> Git hooks keep the tooling managed by mise synchronized after checkouts and merges.

## running

Project tasks are exposed through `mise`.

### `mise run dev`

Runs the project.

### `mise run test`

Runs the test suite.

### `mise run lint`

Runs the configured linters.

### `mise run format`

Formats the project.

### `mise run build`

Builds the project.

## ci

CI is provided by GitHub Actions through [`.github/workflows/checks.yml`](https://github.com/mkvlrn/template-go/blob/main/.github/workflows/checks.yml).

It runs the project's formatting, linting, testing, and build checks.

## license

[MIT](https://github.com/mkvlrn/template-go/blob/main/LICENSE)
