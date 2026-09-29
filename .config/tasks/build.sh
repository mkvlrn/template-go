#!/usr/bin/env bash
#MISE description="Build the Go application binary"

set -euo pipefail

rm -rf ./bin
mkdir -p ./bin
mise exec -- go build -o ./bin/template-go ./cmd/template-go "$@"
