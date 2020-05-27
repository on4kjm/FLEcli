# Enable GO

Notes to enable/configure GO on my Mac

```
# GOlang related
export GOPATH="${HOME}/.go"
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
test -d "${GOPATH}" || mkdir "${GOPATH}"
test -d "${GOPATH}/src/github.com" || mkdir -p "${GOPATH}/src/github.com"
```

* Use Brew to instal Go
* Load the various tools in in VisualCode

* `go mod init example.com/user/hello`
* Structuring a go project: https://github.com/golang-standards/project-layout

## Enable cobra

* `go get -u github.com/spf13/cobra/cobra`
* `go mod init FLEcli`
* `cobra init --pkg-name FLEcli`
* `cobra add load` adds a load command