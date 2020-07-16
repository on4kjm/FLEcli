#!/bin/bash

set -e

./build.sh

mkdir -p test/output/temp

echo "--------------------------"
./FLEcli -i test/data/fle-1.txt load
echo "--------------------------"
./FLEcli -i test/data/ON4KJM@ONFF-025920200524.txt load
echo "--------------------------"
./FLEcli -i test/FLE-sample/sota_wwff.txt csv -o=test/output/temp/sota_wwff.csv --interpolate --overwrite
diff test/output/temp/sota_wwff.csv test/FLE-sample/sota_wwff.csv --strip-trailing-cr && echo "no difference" || echo "differences!"