package main

import (
	"log"
	"net"

	"github.com/ryuku-hisa/gRPC-data-stream-server/server/handler"
	"google.golang.org/grpc"
)

const (
	port = "50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	server := grpc.NewServer()

	handler.NewUploadServer(server)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
