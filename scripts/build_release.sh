#!/usr/bin/env bash

cd ..

platform=("windows" "linux")
arch=("amd64" "386")
for i in "${platform[@]}"; do
    for j in "${arch[@]}"; do
        CGO_ENABLE=true GOOS=$i GOARCH=$j go build -o "dist/release/mul-$i-$j$([ "$i" = "windows" ] && echo ".exe" || true)"
        echo "Building $i $j"
    done
done
