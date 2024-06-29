#!/usr/bin/env bash

cd ..

platform=("linux")
arch=("amd64" "386")
for i in "${platform[@]}"; do
    for j in "${arch[@]}"; do
        CGO_ENABLE=true GOOS=$i GOARCH=$j go build -o "dist/release/mul-$i-$j"
        echo "$i-$j"
    done
done
