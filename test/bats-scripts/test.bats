#!/usr/bin/env bats

@test "Does the version show?" {
    output=$(docker run --rm on4kjm/flecli:latest version -d)
    echo 'status:' $status
    echo 'output:' $output
}