package handlers

import (
	"context"
	"testing"

	"github.com/rigmas/microservices/order/handlers/order_grpc"
	"github.com/rigmas/microservices/product/handlers/product_grpc"
	"github.com/rigmas/microservices/product/mocks"
	"github.com/stretchr/testify/assert"
)

func TestOrderList(t *testing.T) {
	var products []*product_grpc.Product
	product := &product_grpc.Product{
		Id:          "1abc",
		Title:       "Product test",
		Description: "Test",
		Price:       18_000,
		Quantity:    3,
	}

	products = append(products, product)

	productClient := new(mocks.ProductServiceClient)
	reqProduct := &product_grpc.ProductListRequest{}
	productClient.On(
		"ProductList", context.TODO(), reqProduct,
	).Return(&product_grpc.ProductListResponse{Products: products}, nil)

	s := Server{
		Orders:        "",
		ProductClient: productClient,
	}

	reqOrder := &order_grpc.OrderListRequest{}
	srv, err := s.OrderList(context.TODO(), reqOrder)
	assert.NoError(t, err)
	assert.Equal(t, int32(54000), srv.Orders[0].PaymentAmount)
}
