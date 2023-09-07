package main

import (
	"context"
	"log"
	"net"
	pb "github.com/bxvip/grpc-demo/pkg/proto/user"
	"google.golang.org/grpc"
)

const port = ":50051"

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// 模拟用户登录过程，返回登录结果和 token
	return &pb.LoginResponse{
		Success: true,
		Token:   "user_token",
		Message: "User login successfully",
	}, nil
}

func (s *server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	// 模拟用户退出登录过程，返回退出结果
	return &pb.LogoutResponse{
		Success: true,
		Message: "User logout successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Printf("user service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
