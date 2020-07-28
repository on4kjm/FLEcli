#!/bin/bash

set -e

go test ./...

commitRef=$(git rev-parse HEAD)
tag=$(git describe --tags)
version="${tag}_(PrivateBuild)"
buildDate=$(date -u +"%FT%TZ")
go build -ldflags="-s -w -X=FLEcli/flecmd.version=${version} -X=FLEcli/flecmd.commit=${commitRef} -X=FLEcli/flecmd.date=${buildDate} -X=FLEcli/flecmd.builtBy=${USER}"
