#!/bin/bash

start=$SECONDS

echo "Run 'go build' command."
GOOS=linux GOARCH=arm go build -o ./bin/ ./cmd/can-demo/can-demo.go

echo "Run 'snapcraft' commands."
export SNAPCRAFT_BUILD_ENVIRONMENT=host
snapcraft clean
snapcraft

echo "Done compiling in $(($SECONDS - $start)) seconds."
