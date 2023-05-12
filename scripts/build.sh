#!/bin/bash

set -exl
cd `dirname $0`

docker buildx build --platform linux/amd64 -t hexydev/data-bus-receiver:0.0.0 -f ../Dockerfile --load --target=app ../