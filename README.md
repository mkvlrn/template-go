# template-go

![build](https://img.shields.io/github/actions/workflow/status/mkvlrn/template-go/checks.yml?branch=main&style=flat&logo=github)
[![template](https://img.shields.io/badge/template-use_this_template-2ea44f?style=flat&logo=github)](https://github.com/mkvlrn/template-go/generate)
[![mise](https://mise-versions.jdx.dev/badge.svg)](https://mise.jdx.dev)
![license](https://img.shields.io/github/license/mkvlrn/template-go?style=flat)

A sane, opinionated template for Go projects with a consistent dev environment powered by mise.

> [!NOTE]
> This template provides a [mise](https://mise.jdx.dev) configuration to make it easy to keep node and pnpm versions in sync. It handles tool versioning only, while Makefile is the still the task runner.

Uses, among other tools/packages:

- [golangci-lint](https://golangci-lint.run) for linting and formatting
- [lefthook](https://github.com/evilmartians/lefthook) for git hooks
- [cocogitto](https://github.com/cocogitto/cocogitto) for commit message linting

## requirements and dependencies

If you use [mise](https://mise.jdx.dev) and run `mise install` in the project root, you'll have the correct go and other tools installed.

This is _by far_ the easiest way to keep your environment consistent across different machines and team members, no matter the frequency of version updates. I'm not affiliated with mise but I wholeheartedly recommend it, so check it here: https://mise.jdx.dev.

If not using mise, make sure you have:

- go 1.26.1
- golangci-lint 2.11.4
- lefthook 2.1.5
- cocogitto 7.0.0

> [!NOTE]
> Git hooks are in place to make sure both the tooling managed by mise and the project dependencies are synced with each checkout and merge.

## tasks

`make dev <name>`

Runs a program from `./cmd/<name>`.

`make build`

Builds all programs from `./cmd` into `./bin`.

`make test`

Runs all tests.

`make lint`

Runs golangci-lint with the `standard` linter set.

`make format`

Formats code using gofumpt via golangci-lint.

`make format-check`

Checks formatting without modifying files. Exits non-zero if anything needs formatting.

## ci

This repository uses GitHub Actions for CI. The workflow is defined in `.github/workflows/checks.yml`.

## vscode

You might want to install the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go). The `.editorconfig` file handles basic formatting rules.

## license

MIT
