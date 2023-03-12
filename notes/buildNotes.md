# Build Notes

## Bootstrap

Install [Cobra CLI tool](https://github.com/spf13/cobra) `go install github.com/spf13/cobra-cli@latest`

[Install](https://goreleaser.com/install/) [goreleaser](https://goreleaser.com/)

* `go build && go install`

## to build with goreleaser

* `goreleaser --snapshot --skip-publish --rm-dist`
* `dist/FLEcli_darwin_amd64/FLEcli`
* `docker build -f docker/Dockerfile -t test/test dist/FLEcli_linux_arm64`

## to push a release

* on the master branch `git tag v0.1.1`
* `git push --tag`
