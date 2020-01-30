#!/usr/bin/env bash
# Set GOOS environment variable
export GOOS="linux"

go build 

# build DOCKER container
docker build -t lmburu/gateway .

# delete GO executable
go clean

# push the container image to Docker Hub
docker push lmburu/gateway:latest

# Starts web server
ssh ec2-user@ec2-13-59-0-236.us-east-2.compute.amazonaws.com  < deploy.sh