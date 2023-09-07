package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// 获取当前文件夹的绝对路径
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current directory: %v", err)
	}
	log.Println(currentDir)

	// 启动 user 模块的服务
	userCmd := exec.Command("go", "run", filepath.Join(currentDir, "pkg/server/user", "user.go"))
	userCmd.Dir = filepath.Join(currentDir, "pkg/server/user")
	userCmd.Stdout = os.Stdout
	userCmd.Stderr = os.Stderr
	err = userCmd.Start()
	if err != nil {
		log.Fatalf("failed to start user service: %v", err)
	}
	log.Println("user service started")

	// 启动 payment 模块的服务
	paymentCmd := exec.Command("go", "run", filepath.Join(currentDir, "pkg/server/pay", "pay.go"))
	paymentCmd.Dir = filepath.Join(currentDir, "pkg/server/pay")
	err = paymentCmd.Start()
	if err != nil {
		log.Fatalf("failed to start payment service: %v", err)
	}
	log.Println("payment service started")

	// 启动 chat 模块的服务
	chatCmd := exec.Command("go", "run", filepath.Join(currentDir, "pkg/server/chat", "chat.go"))
	chatCmd.Dir = filepath.Join(currentDir, "pkg/server/chat")
	err = chatCmd.Start()
	if err != nil {
		log.Fatalf("failed to start chat service: %v", err)
	}
	log.Println("chat service started")

	// 等待所有服务退出
	err = userCmd.Wait()
	if err != nil {
		log.Fatalf("user service exited with error: %v", err)
	}
	log.Println("user service stopped")

	err = paymentCmd.Wait()
	if err != nil {
		log.Fatalf("payment service exited with error: %v", err)
	}
	log.Println("payment service stopped")

	err = chatCmd.Wait()
	if err != nil {
		log.Fatalf("chat service exited with error: %v", err)
	}
	log.Println("chat service stopped")
}
