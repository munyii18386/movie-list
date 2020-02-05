#!/bin/bash 

ENTRY=$PWD

echo "ACCESSING CLIENT APP..."
cd client/my-app
echo "BUILD/DEPLOY REACH APP..."
sh client-build.sh

echo "EXITING CLIENT APP.."
cd $ENTRY

echo "ACCESSING GATEWAY SERVER..."
cd server/gateway/
echo "BUILD/DEPLOY GATEWAY.."
sh gateway-build.sh





