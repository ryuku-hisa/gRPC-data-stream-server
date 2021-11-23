package handler

import (
  "io"
	"os"
	"path/filepath"

	pb "github.com/ryuku-hisa/gRPC-data-stream-client/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewUploadServer(gserver *grpc.Server) {
	var uploadserver pb.UploadHandlerServer
	pb.RegisterUploadHandlerServer(gserver, uploadserver)
	reflection.Register(gserver)
}

type server struct{}

func (s *server) Upload(stream pb.UploadHandler_UploadServer) error {
  err := os.MkdirAll("Sample", 0777)
	if err != nil {
		return err
	}
	file, err := os.Create(filepath.Join("Sample", "sample.mp4"))
	if err != nil {
		return err
	}
	defer file.Close()

  var int cnt = 0;
  for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		file.Write(resp.VideoData)
    fmt.Println(++cnt)
	}
	err = stream.SendAndClose(&pb.UploadResponse{
		UploadStatus: "OK",
	})
	if err != nil {
		return err
	}
	return nil
}
