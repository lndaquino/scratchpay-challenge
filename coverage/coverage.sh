#!/bin/sh

set -e

[ -z "$DEBUG" ] || set -x

echo "\n===> Running tests...\n"
go test -coverprofile=./coverage/cover.out ./...

echo "\n===> Total coverage: "
go tool cover -func ./coverage/cover.out | grep total | awk '{print $3}'

echo "\n===> Preparing coverage html...\n"
go tool cover -html=./coverage/cover.out -o ./coverage/cover.html

echo "\n===> Open ./coverage/cover.html in web browser to see coverage analysis per file...\n"