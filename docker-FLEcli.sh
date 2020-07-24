#!/bin/bash

CURRENT_UID = $(id -u):$(id -g)
docker run --rm -ti --user $(CURRENT_UID) -v $(pwd):/FLEcli_data jmmeessen/flecli "$@"
