
## goReleaser built Docker images

* `goreleaser --snapshot --skip-publish --rm-dist`
* docker image is named "on4kjm/flecli". It creates two labels automatically ("latest" and the last release)

## Building the image by hand
(this assumes that the Linux executable is available in `dist/`, thus as built by GoReleaser)

* `docker build -f docker/Dockerfile -t test/test dist/FLEcli_linux_amd64` will create a image called `test/test:latest`.

## Running the container

To start and execute the `<FLEcli command>` use : `docker run --rm -i --user $(id -u):$(id -g) -v "$(pwd)":/FLEcli_data on4kjm/flecli <FLEcli command>`. If no command is specified, help is displayed.

This bash script (MAC OS or Linux) will do the trick:

````
#!/bin/bash

CURRENT_UID=$(id -u):$(id -g)
docker run --rm -t --user ${CURRENT_UID} -v "$(pwd)":/FLEcli_data on4kjm/flecli:latest "$@"
````

By creating an alias like here after, this command can be called from everywhere. `alias FLEcli="~/myDir/docker-FLEcli.sh"`. To use it, type `FLEcli version` for example.

Important note: when specifying the path of a file (input or output), it must be relative to the directory the container was started in.

## Running bash in the container

Note, this doesn't work with the SCRATCH image. You need to use the Alpine base image. Anyway, if you want to enter the container, you know what I am talking about.

* `docker run --rm -i -v "$(pwd)":/FLEcli_data --entrypoint /bin/sh  on4kjm/flecli`