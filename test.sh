#!/bin/bash

go build
go test ./...

echo "--------------------------"
./FLEcli -i test/data/fle-1.txt load
echo "--------------------------"
./FLEcli -i test/data/ON4KJM@ONFF-025920200524.txt load
echo "--------------------------"
./FLEcli -i test/data/sample_wwff_sota.txt load