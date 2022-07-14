#!/bin/bash -eux

pushd dp-charts-api
  make test-component
popd
