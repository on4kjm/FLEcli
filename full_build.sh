#!/bin/bash

set -e

go test ./...

mkdir -p executables/windows
mkdir -p executables/macos
mkdir -p executables/linux

#Update version

GOOS=windows GOARCH=386 go build -o executables/windows/FLEcli.exe
GOOS=darwin GOARCH=amd64 go build -o executables/macos/FLEcli
GOOS=linux GOARCH=amd64 go build -o executables/linux/FLEcli-amd64

