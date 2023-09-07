package main

import (
	"io"
	"log"
	"net"

	pb "github.com/bxvip/grpc-demo/pkg/proto/chat"
	"google.golang.org/grpc"
)

const port = ":50053"

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) Chat(stream pb.ChatService_ChatServer) error {
	// 模拟双向流的聊天过程，客户端和服务端都可以发送和接收消息
	for {
		req, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		log.Printf("received message from %s: %s", req.Username, req.Message)
		res := &pb.ChatResponse{
			Username: req.Username,
			Message:  req.Message,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	log.Printf("chat service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
