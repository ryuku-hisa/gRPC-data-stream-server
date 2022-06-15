package main

import (
	"path/filepath"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	pb "github.com/ryuku-hisa/gRPC-data-stream-server/proto"
//	"github.com/ryuku-hisa/gRPC-data-stream-server/server/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedUploadHandlerServer
}

func (s *server) Upload(stream pb.UploadHandler_UploadServer) error {
	fmt.Printf("\nRecieving...\n")
  err := os.MkdirAll("Sample", 0777)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath.Join("Sample", "sample.mp4"))
	if err != nil {
		return err
	}
	defer file.Close()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		file.Write(resp.VideoData)
	}
	err = stream.SendAndClose(&pb.UploadResponse{
		UploadStatus: "OK",
	})
	if err != nil {
		return err
	}
  fmt.Println("DONE")
	return nil
}

func main() {
	fmt.Printf("Preparing...\n\n")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("faild to listen: %v", err)
	}
	gserver := grpc.NewServer()

	fmt.Println("--- Recieving ---")
	//handler.NewUploadServer(server)

	uploadserver := &server{}
	pb.RegisterUploadHandlerServer(gserver, uploadserver)
	reflection.Register(gserver)

	if err := gserver.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
