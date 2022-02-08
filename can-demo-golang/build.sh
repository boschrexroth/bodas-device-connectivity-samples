#!/bin/bash
# Copyright (c) 2022 Bosch Rexroth AG
# All rights reserved. See LICENSE file for details.

start=$SECONDS

echo "Run 'go build' command."
GOOS=linux GOARCH=arm go build -o ./bin/ ./cmd/can-demo/can-demo.go

echo "Run 'snapcraft' commands."
export SNAPCRAFT_BUILD_ENVIRONMENT=host
snapcraft clean
snapcraft
snapcraft clean
echo "Done compiling in $(($SECONDS - $start)) seconds."
