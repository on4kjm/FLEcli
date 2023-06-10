# Enable GO

Notes to enable/configure GO on my Mac

```sh
# GOlang related
export GOPATH="${HOME}/.go"
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
test -d "${GOPATH}" || mkdir "${GOPATH}"
test -d "${GOPATH}/src/github.com" || mkdir -p "${GOPATH}/src/github.com"
```

* Use Brew to install Go
* Load the various tools in in VisualCode

* `go mod init example.com/user/hello`
* Structuring a go project: <https://github.com/golang-standards/project-layout>

## Cobra

Cobra documentation:

* [main README](https://github.com/spf13/cobra/blob/main/README.md)
* [Cobra User Guide](https://github.com/spf13/cobra/blob/main/user_guide.md)
* [Cobra CLI](https://github.com/spf13/cobra-cli/blob/main/README.md)

Setting up project to use Cobra:

* `go get -u github.com/spf13/cobra/cobra`
* `go mod init FLEcli`
* `cobra init --pkg-name FLEcli -l MIT`
* `cobra add load` adds a load command

## Reading and processing file

* [Go by Example: Reading Files](https://gobyexample.com/reading-files)
* <https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go>
* <https://www.golangprograms.com/golang-read-file-line-by-line-to-string.html>

## Regex

* useful site: <https://regex101.com>
* [Regex to verify Hamradio callsign and retrieve prefix](https://regex101.com/library/6QhGuD)
* [Hamradio callsign](https://regex101.com/library/uP6xD2)
* [Verify amateur radio call sign meets ITU format](https://regex101.com/library/gS6qG8)

## Parsing 2

* [Effective text parsing in golang](https://medium.com/@TobiasSchmidt89/effective-text-parsing-in-golang-163d13784288)

## Amateur related Golang libray

* <https://github.com/tzneal/ham-go>
