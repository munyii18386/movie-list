#!/bin/bash
docker build -t lmburu/my-app .

docker push lmburu/my-app

ssh -T ec2-user@ec2-13-59-0-236.us-east-2.compute.amazonaws.com  < deploy.sh

