package main

import (
	"context"
	"log"
	"net"

	pb "github.com/bxvip/grpc-demo/pkg/proto/pay"
	"google.golang.org/grpc"
)

const port = ":50052"

type server struct {
	pb.UnimplementedPaymentServiceServer
}

func (s *server) GetPaymentInfo(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentInfo, error) {
	// 模拟从数据库中获取订单的支付信息，返回一个 PaymentInfo 结构体
	return &pb.PaymentInfo{
		UserId:      req.UserId,
		OrderId:     req.OrderId,
		Amount:      100.00,
		Description: "Order payment",
	}, nil
}

func (s *server) Pay(ctx context.Context, req *pb.PaymentInfo) (*pb.PaymentResult, error) {
	// 模拟调用第三方支付接口进行支付，返回支付结果
	return &pb.PaymentResult{
		Success: true,
		Message: "Payment succeed",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, &server{})
	log.Printf("payment service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
