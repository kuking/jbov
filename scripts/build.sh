#!/bin/bash

set -eo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
OUT_DIR="$DIR/../out"
mkdir -p $OUT_DIR

go build \
  -o $OUT_DIR/jbov \
  -ldflags "-X github.com/kuking/jbov/Build=`git rev-parse --short HEAD`" \
  github.com/kuking/jbov/cmd/jbov
