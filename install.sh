#!/usr/bin/env bash

git pull origin master
cd ../docker && docker-compose stop && docker-compose rm --all --force
docker stop $(docker ps -a -q --filter ancestor=zhanat87/golang-grpc-protobuf-server)
docker rm $(docker ps -a -q --filter ancestor=zhanat87/golang-grpc-protobuf-server) -f
docker rmi $(docker images --filter=reference='zhanat87/golang-grpc-protobuf-server') -f
docker pull zhanat87/golang-grpc-protobuf-server
docker-compose up -d
docker images

echo "install success"