#!/bin/bash

mkdir -p build

echo "Generating builds..."
GOOS=windows GOARCH=amd64 go build -o build/app-windows.exe
GOOS=linux GOARCH=amd64 go build -o build/app-linux
GOOS=darwin GOARCH=amd64 go build -o build/app-macos
GOOS=darwin GOARCH=arm64 go build -o build/app-macos-arm

echo "Successfuly created at /build"