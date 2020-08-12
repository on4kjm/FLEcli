#!/usr/bin/env bats

# the file is assumed to be run from the root of the project (test data location)

@test "Does the version show?" {
    output=$(docker run --rm on4kjm/flecli:latest version -d)
    echo 'status:' $status
    echo 'output:' $output
}


@test "Can a simple file be loaded?" {
    output=$(test/docker-FLEcli.sh load -i test/data/fle-1.txt)
    echo 'status:' $status
    echo 'output:' $output
}


@test "Can a more complex file be loaded (docker)?" {
    output=$(test/docker-FLEcli.sh load -i test/data/ON4KJM@ONFF-025920200524.txt)
    echo 'status:' $status
    echo 'output:' $output
}

@test "Is the generated SOTA csv equivalent to the canonical one" {
    mkdir -p test/output/temp
    output=$(test/docker-FLEcli.sh csv -o -i test/FLE-sample/sota_wwff.txt test/output/temp/sota_wwff.csv)
    echo 'status:' $status
    echo 'output:' $output
    diff test/output/temp/sota_wwff.csv test/FLE-sample/sota_wwff.csv --strip-trailing-cr
}


@test "Direct processing a FLE file with no QSO must fail" {
    run ./FLEcli csv -o -i test/data/fle-4-no-qso.txt
    [ "$status" -eq 1 ]
}

@test "Processing a FLE file with no QSO must fail" {
    run test/docker-FLEcli.sh csv -o -i test/data/fle-4-no-qso.txt
    [ "$status" -eq 1 ]
}

