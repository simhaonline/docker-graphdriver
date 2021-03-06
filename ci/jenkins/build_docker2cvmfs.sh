#!/bin/bash

set -e

die() {
  echo "$1"
  exit 1
}

[ "x$CVMFS_SOURCE_LOCATION" != x ] || die "CVMFS_SOURCE_LOCATION missing"
[ "x$CVMFS_BUILD_LOCATION" != x ] || die "CVMFS_BUILD_LOCATION missing"

export GOPATH="$CVMFS_SOURCE_LOCATION/.."
D2C_ROOT="github.com/cvmfs/docker-graphdriver/docker2cvmfs"
GIT_COMMIT=$(cd $CVMFS_SOURCE_LOCATION/$D2C_ROOT && git rev-parse HEAD)

cd $CVMFS_BUILD_LOCATION
echo "Building docker2cvmfs"

VERSION=$(cat $CVMFS_SOURCE_LOCATION/$D2C_ROOT/VERSION)
mkdir -p $VERSION
pushd $VERSION
go build \
  -ldflags="-X main.version=$VERSION -X main.git_hash=$GIT_COMMIT" \
  -v "$D2C_ROOT"
popd

