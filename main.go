package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"github.com/Zhanat87/golang-grpc-protobuf-server/server"
	grpc_local "github.com/Zhanat87/go/grpc"
)

const (
	port = ":50051"
)

func main() {
	// lsof -i:50051
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	grpc_local.RegisterGrpcServiceServer(s, &server.Server{})
	s.Serve(lis)
}