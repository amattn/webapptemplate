#!/bin/sh

# if one of our commands returns an error, stop execution of this script
set -o errexit 


# build on the native or default platform
echo "building native platform"
go build

# test on the native or default platform
echo "testing native platform"
go test

# I like gox as a cross compilation tool: https://github.com/mitchellh/gox
# install with:
# go get github.com/mitchellh/gox

# build for our test or deployment platforms
# normally we do local development on darwin/amd64 and deploy to linux/amd64.
# feel free to add or remove if your situation differs
# also, in the normal case, most of the output of gox is redundant with
# the output from go build above, so in the normal case, we just 
# redirect to a build log

GOX_BUILD_LOG="gox_build.log"
date >> $GOX_BUILD_LOG
echo "crosscompiling: gox -osarch=\"darwin/amd64\" -osarch=\"linux/amd64\""
if ! gox -osarch="darwin/amd64" -osarch="linux/amd64" >> $GOX_BUILD_LOG 2>&1; then
    echo "FAILURE: gox command failed to build for deployment architecture"
    echo "check gox_build.log for more info"
fi

