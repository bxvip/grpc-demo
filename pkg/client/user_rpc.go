package client

import (
	userpb "github.com/bxvip/grpc-demo/pkg/proto/user"
)

type UserRpcClient struct {
	userClient userpb.UserServiceClient
}

func NewUserRpcClient() *UserRpcClient {
	return &UserRpcClient{

	}
}
