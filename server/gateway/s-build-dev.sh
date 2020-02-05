#!/usr/bin/env bash
export GOOS="linux"

go build 

# build DOCKER container
docker build -t lmburu/gateway .

# delete GO executable
go clean

# push the container image to Docker Hub
docker push lmburu/gateway

# Starts web server
sh s-deploy-dev.sh