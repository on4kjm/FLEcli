#!/bin/bash

echo "# Usage"  > help.txt
echo " " >> help.txt
echo " " >> help.txt

echo "## Overview"  >> help.txt
echo "\`\`\`"  >> help.txt
../dist/FLEcli_darwin_amd64/FLEcli >> help.txt
echo "\`\`\`"  >> help.txt
echo " " >> help.txt
echo " " >> help.txt

echo "## \"LOAD\" command"  >> help.txt
echo "\`\`\`"  >> help.txt
../dist/FLEcli_darwin_amd64/FLEcli load --help >> help.txt
echo "\`\`\`"  >> help.txt
echo " " >> help.txt
echo " " >> help.txt

echo "## \"ADIF\" command"  >> help.txt
echo "\`\`\`"  >> help.txt
../dist/FLEcli_darwin_amd64/FLEcli adif --help >> help.txt
echo "\`\`\`"  >> help.txt
echo " " >> help.txt
echo " " >> help.txt

echo "## \"CSV\" command"  >> help.txt
echo "\`\`\`"  >> help.txt
../dist/FLEcli_darwin_amd64/FLEcli csv --help >> help.txt
echo "\`\`\`"  >> help.txt
echo " " >> help.txt
echo " " >> help.txt

echo "## \"VERSION\" command"  >> help.txt
echo "\`\`\`"  >> help.txt
../dist/FLEcli_darwin_amd64/FLEcli version --help >> help.txt
echo "\`\`\`"  >> help.txt

echo "The normal output looks like \`FLEcli version: v0.1.2\`. The detailled output gives additionaly the Git commit hash. the date and time of build and who built the release."   >> help.txt