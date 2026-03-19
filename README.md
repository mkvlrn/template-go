# template-go

![build](https://img.shields.io/github/actions/workflow/status/mkvlrn/template-go/checks.yml?branch=main&style=flat&logo=github)
[![template](https://img.shields.io/badge/template-use_this_template-2ea44f?style=flat&logo=github)](https://github.com/mkvlrn/template-go/generate)
[![mise](https://badgen.net/badge/mise/en-place/1c1614?icon=https://raw.githubusercontent.com/mkvlrn/template-node/refs/heads/main/.github/assets/mise-logo-dark.svg)](https://mise.jdx.dev)
![license](https://img.shields.io/github/license/mkvlrn/template-go?style=flat)

A sane, opinionated template for Go projects with a consistent dev environment powered by mise.

> [!IMPORTANT]
> This template requires **mise**. It manages the correct versions of runtimes and tooling, such as Node itself, pnpm, and others.
>
> Check https://mise.jdx.dev

mise handles tool versioning and task running, replacing both version managers and Makefiles with a single, cross-platform tool.

## uses

- [golangci-lint](https://golangci-lint.run) for linting and formatting
- [lefthook](https://github.com/evilmartians/lefthook) for git hooks
- [commitlint](https://github.com/conventionalcommit/commitlint) for linting commit messages

## setup

To ensure a reproducible environment, [mise](https://mise.jdx.dev/) is used:

1. **Install mise**: https://mise.jdx.dev/getting-started.html#installing-mise-cli
2. **Activate mise**: https://mise.jdx.dev/getting-started.html#activate-mise
3. **Run setup**:
   ```bash
   mise setup
   ```

This task trusts the project config, installs CLI tools (Node, pnpm, ncu), and runs pnpm install. All other scripts are standard package.json commands.

> [!NOTE]
> Git hooks are in place to make sure both the tooling managed by mise and the project dependencies are synced with each checkout and merge.

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

## license

MIT
