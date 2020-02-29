#!/bin/bash

docker rm -f wom-redis

docker run -d --name wom-redis --network bucket redis