#!/bin/bash
#docker network create bucket (run once)

#docker rm -f lmburu/my-app

docker rm -f wom

docker run -d \
--name wom \
-p 5000:5000 \
--network bucket \
-e PORT=5000 \
-e NAME=wom \
lmburu/my-app 