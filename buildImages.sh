#!/bin/bash
docker build -t hectorgabucio/random-micro --build-arg project=./cmd/random-micro/ .
docker build -t hectorgabucio/backend --build-arg project=./cmd/backend/ .

docker push hectorgabucio/random-micro
docker push hectorgabucio/backend