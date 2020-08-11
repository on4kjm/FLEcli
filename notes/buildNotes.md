## to build with goreleaser

* `goreleaser --snapshot --skip-publish --rm-dist`
* `dist/FLEcli_darwin_amd64/FLEcli`
* `docker build -f docker/Dockerfile -t test/test dist/FLEcli_linux_arm64`

## to push a release

* on the master branch `git tag v0.1.1`
* `git push --tag`