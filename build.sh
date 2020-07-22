#!/bin/bash

set -e

go test ./...

commitRef=$(git rev-parse HEAD)
tag=$(git describe --tags)
version="${tag}_(PrivateBuild)"
buildDate=$(date -u +"%FT%TZ")
go build -ldflags="-s -w -X=FLEcli/cmd.version=${version} -X=FLEcli/cmd.commit=${commitRef} -X=FLEcli/cmd.date=${buildDate} -X=FLEcli/cmd.builtBy=${USER}"
