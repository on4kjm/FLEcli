#!/bin/bash

../dist/FLEcli_darwin_amd64/FLEcli > help.txt
echo "----" >> help.txt
../dist/FLEcli_darwin_amd64/FLEcli load --help >> help.txt
echo "----" >> help.txt
../dist/FLEcli_darwin_amd64/FLEcli adif --help >> help.txt
echo "----" >> help.txt
../dist/FLEcli_darwin_amd64/FLEcli csv --help >> help.txt
echo "----" >> help.txt
../dist/FLEcli_darwin_amd64/FLEcli version --help >> help.txt