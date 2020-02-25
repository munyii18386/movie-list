#!/bin/bash

# export MYSQL_ROOT_PASSWORD=supersecret

docker rm -f wom-database


docker run -d \
-p 3306:3306 \
--network bucket \
--name wom-database \
-e MYSQL_ROOT_PASSWORD=supersecret \
-e MYSQL_ROOT_HOST=% \
-e MYSQL_DATABASE=movie-list-database \
lmburu/database