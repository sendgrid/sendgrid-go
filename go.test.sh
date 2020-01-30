#!/usr/bin/env bash

set -e

for d in $(go list ./... | grep -v -E '/vendor|/examples|/docker'); do
    go test -race "$d"
done
