FROM alpine:latest
MAINTAINER Iskakov Zhanat <iskakov_zhanat@mail.ru>
ADD golang-grpc-protobuf-server /usr/bin/golang-grpc-protobuf-server
ENTRYPOINT ["golang-grpc-protobuf-server"]
EXPOSE 50051