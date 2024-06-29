#!/usr/bin/env bash

export GIN_MODE=debug

cd ..
CGO_ENABLE=true go build -o dist/debug/mul-server

cd dist/debug || exit
./mul-server
