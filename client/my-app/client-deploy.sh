#!/bin/bash
docker rm -f lmburu/my-app

docker pull lmburu/my-app

docker rm -f wom

docker run -d \
--name wom \
-p 5000:5000 \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
--network bucket \
-e PORT=5000 \
-e NAME=wom \
lmburu/my-app 