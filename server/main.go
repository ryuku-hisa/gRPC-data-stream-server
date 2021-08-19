package main

import (
	"log"
	"net"

	pb "github.com/Ryuku-Hisa/gRPC-data-stream-server/proto"
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

	pb.NewUploadServer(server)
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
