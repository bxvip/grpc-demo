package client

import (
	"fmt"
	"sync"
)

type GrpcManager struct {
	camera *CameraManager
}

var (
	once   sync.Once
	camera *CameraManager
)

func loadCamera() {
	// 在这个函数中进行资源的初始化和加载操作
	fmt.Println("Loading data...")
	camera = NewCameraManager()
}

func (m *GrpcManager) GetCamera() *CameraManager {
	once.Do(loadCamera)
	return camera
}
