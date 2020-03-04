#!/bin/bash
docker rm -f lmburu/gateway

docker pull lmburu/gateway

docker rm -f wom-api

export TLSCERT=/etc/letsencrypt/live/oddgarden.net/fullchain.pem
export TLSKEY=/etc/letsencrypt/live/oddgarden.net/privkey.pem
export DSN="root:supersecret@tcp(wom-database)/movie-list-database"
export REDISADDR="wom-redis:6379"
export SESSIONKEY="sessionkey"

docker run -d --name wom-api \
-p 443:443 \
--network bucket \
-v /etc/letsencrypt:/etc/letsencrypt:ro \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-e DSN=$DSN \
-e REDISADDR=$REDISADDR \
-e SESSIONKEY=$SESSIONKEY \
-e REACT=wom:5000 \
lmburu/gateway 