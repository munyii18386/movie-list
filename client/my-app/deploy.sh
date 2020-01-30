#!/bin/bash
docker rm -f lmburu/my-app

docker pull lmburu/my-app

docker rm -f wom

docker run -d --name wom -p 80:5000 -v /etc/letsencrypt:/etc/letsencrypt:ro lmburu/my-app 