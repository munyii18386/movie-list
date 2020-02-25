#!/bin/bash

docker rm -f wom-api


export TLSCERT=/certs/server.crt
export TLSKEY=/certs/server.key
export DSN="root:supersecret@tcp(wom-database)/movie-list-database"

docker run -d --name wom-api \
-p 443:443 \
--network bucket \
-e REACT=wom:5000 \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
-e DSN=$DSN \
-v /Users/lilibug/go/src/movie-list/server/gateway/certs:/certs:ro \
lmburu/gateway

