#!/usr/bin/env bash
#MISE description="Run the Go application"

set -euo pipefail

mise exec -- go run . "$@"
