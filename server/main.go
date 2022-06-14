package main

import (
  "fmt"
  "log"
	"net"

	"github.com/ryuku-hisa/gRPC-data-stream-server/server/handler"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
  fmt.Printf("Preparing...\n\n")

  lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	server := grpc.NewServer()

  fmt.Println("--- Recieving ---")
	handler.NewUploadServer(server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
