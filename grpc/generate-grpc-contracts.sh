#!/bin/env bash
set -e
set +x
set -o pipefail

# CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
# protoc --proto_path=$CURRENT_DIR --go-grpc_out=. $CURRENT_DIR/**/*.proto

cd "$( dirname "${BASH_SOURCE[0]}" )"
# protoc --go-grpc_out=. ./**/*.proto
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./**/*.proto
