#!/bin/bash

set -e

gf build main.go -a amd64 -s linux -p ./temp
gf docker main.go -p -t weikedata/auditlimit:latest
now=$(date +"%Y%m%d%H%M%S")
# 以当前时间为版本号
docker tag weikedata/auditlimit:latest weikedata/auditlimit:$now
docker push weikedata/auditlimit:$now
echo "release success" $now
