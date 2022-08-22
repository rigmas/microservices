package handlers

import (
	"context"
	"testing"

	"github.com/rigmas/microservices/customer/handlers/customer_grpc"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	s := Server{}

	mockReq := &customer_grpc.RegisterRequest{
		Username: "acb123",
		Password: "xxx123",
	}

	srv, err := s.Register(context.TODO(), mockReq)
	assert.NoError(t, err)
	assert.Equal(t, "username:acb123-password:xxx123", srv.Token)
}
