#!/bin/bash

set -e

VERSION_FROM_GIT=$(git describe --tags)
BUILD_TIME=$(date +%F)

echo "Building version ${VERSION_FROM_GIT}"

echo "package cmd" > cmd/FLEcli_version.go
echo " " >> cmd/FLEcli_version.go
echo "//VersionString is the version that will be displayed with the -v switch" >> cmd/FLEcli_version.go
echo "const VersionString = \"${VERSION_FROM_GIT} (${BUILD_TIME})\"" >> cmd/FLEcli_version.go