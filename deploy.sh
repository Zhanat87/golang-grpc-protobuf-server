#!/usr/bin/env bash

CGO_ENABLED=0 GOOS=linux go build
git add . && git commit -m 'deploy' && git push origin master
# stop & remove all docker containers
docker stop $(docker ps -a -q)
# remove image
docker rmi $(docker images --filter=reference='zhanat87/golang-grpc-protobuf-server') -f
# create new docker image, push to docker hub and pull
docker build -t zhanat87/golang-grpc-protobuf-server .
docker push zhanat87/golang-grpc-protobuf-server
# list of all docker images on host machine
docker images

echo "deploy success"