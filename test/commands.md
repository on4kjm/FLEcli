# Commands for testing

* `./FLEcli -i test/data/fle-1.txt load`
* `./FLEcli -i test/data/sample_contest_ru.txt load`
* `go test ./...` runs the unit tests, from project root

* `./FLEcli -i test/data/ON4KJM@ONFF-025920200524.txt --interpolate adif --wwff --overwrite`
* `./FLEcli adif -i=test/data/ON4KJM@ONFF-025920200524.txt --interpolate --wwff --overwrite`

Install Bats: `brew install bats-core`
