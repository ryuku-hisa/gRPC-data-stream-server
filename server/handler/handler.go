package handler

import (
  "io"
	"os"
	"path/filepath"
  "fmt"

	pb "github.com/ryuku-hisa/gRPC-data-stream-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{
  pb.UnimplementedGreeterServer
}

func NewUploadServer(gserver *grpc.Server) {
//  uploadserver := &server{}
//  pb.RegisterUploadHandlerServer(gserver, uploadserver)
//  reflection.Register(gserver)
}

func (s *server) Upload(stream pb.UploadHandler_UploadServer) error {
  fmt.Println("called Upload")
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
	return nil
}
