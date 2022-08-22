package handlers

import (
	"context"
	"testing"

	"github.com/rigmas/microservices/product/handlers/product_grpc"
	"github.com/stretchr/testify/assert"
)

func TestProductList(t *testing.T) {
	s := Server{}

	mockReq := &product_grpc.ProductListRequest{}
	srv, err := s.ProductList(context.TODO(), mockReq)
	assert.NoError(t, err)
	assert.Equal(t, "Iphone", srv.Products[0].Title)
	assert.Equal(t, "Airpods Pro 2", srv.Products[1].Description)
}
