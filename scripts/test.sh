#!/usr/bin/env bash

rm coverage.out test-result.out
set -e
echo "mode: atomic" > coverage.out
touch test-result.out

if [[ $1 ]]; then
        VERBOSE="-v"
else
        VERBOSE=""
fi

for d in $(go list ./... | grep -v /vendor/ | grep -v /offerbiddingservrpc ); do
    COVER_PKG=$d

        go test $VERBOSE --race -coverprofile=profile.out -coverpkg=$COVER_PKG -covermode=atomic $d | tee result.out
        if [ -f profile.out ]; then
                tail -n +2 profile.out >> coverage.out
                rm profile.out
        fi
        if [ -f result.out ]; then
                tail -n +1 result.out >> test-result.out
                rm result.out
        fi
done
