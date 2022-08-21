package handlers

import (
	"context"
	"fmt"

	"github.com/rigmas/microservices/customer/handlers/customer_grpc"
)

type Server struct {
	Token string

	customer_grpc.UnimplementedCustomerServiceServer
}

func (s *Server) Register(_ context.Context,
	req *customer_grpc.RegisterRequest) (*customer_grpc.RegisterResponse, error) {

	username := req.Username
	password := req.Password

	return &customer_grpc.RegisterResponse{
		Token: fmt.Sprintf("username:%v-password:%v", username, password),
	}, nil
}
