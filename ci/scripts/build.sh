#!/bin/bash -eux

pushd dp-charts-api
  make build
  cp build/dp-charts-api Dockerfile.concourse ../build
popd
