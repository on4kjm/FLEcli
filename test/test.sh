#!/bin/bash

set -e

./build.sh



echo "--------------------------"
./FLEcli load -i test/data/fle-1.txt
echo "--------------------------"
./FLEcli load -i test/data/ON4KJM@ONFF-025920200524.txt
echo "--------------------------"
mkdir -p test/output/temp
./FLEcli csv -o -i test/FLE-sample/sota_wwff.txt test/output/temp/sota_wwff.csv
diff test/output/temp/sota_wwff.csv test/FLE-sample/sota_wwff.csv --strip-trailing-cr && echo "no difference" || echo "differences!"
echo "--------------------------"
./FLEcli -i test/FLE-sample/sota_wwff.txt adif -o=test/output/temp/sota_wwff.adif --interpolate --overwrite --wwff --sota
