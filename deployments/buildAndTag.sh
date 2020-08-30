#!/bin/bash

docker build \
    -f ./build/Dockerfile \
    --tag $1 \
    --build-arg project=$2 \
    .