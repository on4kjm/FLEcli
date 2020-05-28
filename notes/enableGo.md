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

Cobra documentation available at https://github.com/spf13/cobra/blob/master/cobra/README.md

* `go get -u github.com/spf13/cobra/cobra`
* `go mod init FLEcli`
* `cobra init --pkg-name FLEcli -l MIT`
* `cobra add load` adds a load command

## Reading and processing file 

* [Go by Example: Reading Files](https://gobyexample.com/reading-files)
* https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
