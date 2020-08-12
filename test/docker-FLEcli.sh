#!/bin/bash

CURRENT_UID=$(id -u):$(id -g)
docker run --rm -t --user ${CURRENT_UID} -v $(pwd):/FLEcli_data on4kjm/flecli:latest "$@"
