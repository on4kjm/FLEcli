#!/bin/bash

set -e

./update-version.sh

go build
go test ./...