## Building docker image

* `docker build -t fle_cli .`

## Running the container

* `docker run --rm -i fle_cli version -d`
* `docker run --rm -i -v "$(pwd)":/FLEcli_data fle_cli version -d`

## Running bash in the container

* `docker run --rm -i -v "$(pwd)":/FLEcli_data --entrypoint /bin/sh  fle_cli`