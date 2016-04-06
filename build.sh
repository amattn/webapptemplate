#!/bin/sh

# if one of our commands returns an error, stop execution of this script
set -o errexit 


# build on the native or default platform
go build

# test on the native or default platform
go test

# build for our test or deployment platforms
# normally we do local development on darwin/amd64 and deploy to linux/amd64.
# feel free to add or remove if your situation differs
gox -osarch="darwin/amd64" -osarch="linux/amd64"
