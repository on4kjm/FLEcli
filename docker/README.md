## Building docker image by hand

* `docker build -t fle_cli .`

## goReleaser built Docker images

* `goreleaser --snapshot --skip-publish --rm-dist`
* docker image is named "jmmeessen/flecli". It creates two labels automatically ("latest" and the last release)

## Running the container

* `docker run --rm -i fle_cli version -d`
* `docker run --rm -i -v "$(pwd)":/FLEcli_data fle_cli version -d`
* `docker run --rm -i jmmeessen/flecli version -d`

* `docker run --rm -i --user $(id -u):$(id -g) -v $(pwd):/FLEcli_data jmmeessen/flecli`. 
* `alias FLEcli="docker run --rm  --user $(id -u):$(id -g) -v $(pwd):/FLEcli_data jmmeessen/flecli"`

## Running bash in the container

* `docker run --rm -i -v "$(pwd)":/FLEcli_data --entrypoint /bin/sh  fle_cli`