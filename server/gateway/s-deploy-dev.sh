#!/bin/bash

docker rm -f wom-api

openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj "/CN=localhost" -keyout privkey.pem -out fullchain.pem

export TLSCERT=/Users/lilibug/go/src/movie-list/server/gateway/fullchain.pem
export TLSKEY=/Users/lilibug/go/src/movie-list/server/gateway/privkey.pem

docker run -d --name wom-api \
-p 443:443 \
-e TLSCERT=$TLSCERT \
-e TLSKEY=$TLSKEY \
--network bucket \
-e REACT=wom:5000 \
lmburu/gateway 