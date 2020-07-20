#!/bin/bash

set -e

go test ./...
go build -ldflags=-X=FLEcli/cmd.VersionString=privateBuild
