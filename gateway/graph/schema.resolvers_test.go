package graph

import (
	"context"
	"testing"

	"github.com/rigmas/microservices/customer/handlers/customer_grpc"
	mockCustomer "github.com/rigmas/microservices/customer/mocks"
	"github.com/rigmas/microservices/order/handlers/order_grpc"
	mockOrder "github.com/rigmas/microservices/order/mocks"
	"github.com/rigmas/microservices/product/handlers/product_grpc"
	mockProduct "github.com/rigmas/microservices/product/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	resolver := Resolver{}

	srv, err := resolver.Mutation().Health(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, "ready", srv.Message)
	assert.Equal(t, "healthy", srv.Healthiness)
}

func TestRegister(t *testing.T) {
	customerClient := new(mockCustomer.CustomerServiceClient)
	resolver := Resolver{
		CustomerService: customerClient,
	}

	mockReq := &customer_grpc.RegisterRequest{
		Username: "test_username",
		Password: "test_pass",
	}

	customerClient.On(
		"Register",
		context.TODO(),
		mockReq,
	).Return(&customer_grpc.RegisterResponse{Token: "username:test_username-password:test_pass"}, nil)

	srv, err := resolver.Mutation().Register(context.TODO(), mockReq.Username, mockReq.Password)
	assert.NoError(t, err)
	assert.Equal(t, "username:test_username-password:test_pass", srv.Token)
}

func TestPing(t *testing.T) {
	resolver := Resolver{}

	srv, err := resolver.Query().Ping(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, "pong", srv)
}

func TestProducts(t *testing.T) {
	productClient := new(mockProduct.ProductServiceClient)
	resolver := Resolver{
		ProductService: productClient,
	}

	mockReq := &product_grpc.ProductListRequest{}

	mockProduct := &product_grpc.Product{
		Id:          "1",
		Title:       "Test",
		Description: "mocktestproduct",
		Price:       1_000_000,
		Quantity:    77,
		CreatedAt:   "2022-08-15 17:01:36.399357+03",
	}

	var mockProducts []*product_grpc.Product
	mockProducts = append(mockProducts, mockProduct)

	productClient.On(
		"ProductList",
		context.TODO(),
		mockReq,
	).Return(&product_grpc.ProductListResponse{Products: mockProducts}, nil)

	srv, err := resolver.Query().Products(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, "mocktestproduct", srv[0].Description)
	assert.Equal(t, 77, srv[0].Quantity)
}

func TestOrder(t *testing.T) {
	orderClient := new(mockOrder.OrderServiceClient)
	resolver := Resolver{
		OrderService: orderClient,
	}

	mockReq := &order_grpc.OrderListRequest{UserId: "1"}

	mockOrder := &order_grpc.Order{
		InvoiceNumber: "test123",
		PaymentAmount: 888_000,
		CreatedAt:     "2022-08-15 17:01:36.399357+03",
	}

	mockOrders := make([]*order_grpc.Order, 0)
	mockOrders = append(mockOrders, mockOrder)

	orderClient.On(
		"OrderList",
		context.TODO(),
		mockReq,
	).Return(&order_grpc.OrderListResponse{Orders: mockOrders}, nil)

	srv, err := resolver.Query().Orders(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, "test123", srv[0].InvoiceNumber)
	assert.Equal(t, 888000, srv[0].PaymentAmount)
}
