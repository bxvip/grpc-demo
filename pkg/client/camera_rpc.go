package client

import (
	cameraPb "github.com/bxvip/grpc-demo/pkg/proto/camera"
	"google.golang.org/grpc"
	"log"
)

const (
	port = "50051"
)

type CameraManager struct {
	camera cameraPb.CameraServiceClient
}

func NewCameraManager() *CameraManager {
	camera = &CameraManager{}
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := cameraPb.NewCameraServiceClient(conn)
	camera.camera = c

	return camera
}
