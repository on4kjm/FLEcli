#!/usr/bin/env bats

# the file is assumed to be run from the root of the project (test data location)

@test "Does the version show?" {
    output=$(docker run --rm on4kjm/flecli:latest version -d)
#   [ "$status" -eq 0 ]
}


@test "Can a simple file be loaded?" {
    output=$(test/docker-FLEcli.sh load -i test/data/fle-1.txt)
}


@test "Can a more complex file be loaded?" {
    output=$(test/docker-FLEcli.sh load -i test/data/ON4KJM@ONFF-025920200524.txt)
}

@test "Is the generated SOTA csv equivalent to the canonical one?" {
    mkdir -p test/output/temp
    output=$(test/docker-FLEcli.sh csv -o -i test/FLE-sample/sota_wwff.txt test/output/temp/sota_wwff.csv)
    diff test/output/temp/sota_wwff.csv test/FLE-sample/sota_wwff.csv --strip-trailing-cr
}

@test "Is the generated POTA adif equivalent to the canonical one?" {
    mkdir -p test/output/temp
    output=$(test/docker-FLEcli.sh adif -o -i --pota test/data/sample_pota.txt test/output/temp/sample_pota.adif)
    diff test/output/temp/sample_pota.adif test/output/POTA/sample_pota.adif --strip-trailing-cr
}

@test "Processing a FLE file with no QSO must fail!" {
    run test/docker-FLEcli.sh csv -o -i test/data/fle-4-no-qso.txt
    [ "$status" -eq 1 ]
}

@test "Processing a FLE file with parsing errors must fail!" {
    run test/docker-FLEcli.sh csv -o -i test/data/fle-5-wrong-call.txt
    [ "$status" -eq 1 ]
}

@test "Processing a big FLE file" {
    run test/docker-FLEcli.sh csv -o -i test/data/fle-6-bigFile.txt test/output/temp/fle-6-bigFile.csv
}