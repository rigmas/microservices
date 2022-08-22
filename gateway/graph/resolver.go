package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/rigmas/microservices/customer/handlers/customer_grpc"
	"github.com/rigmas/microservices/order/handlers/order_grpc"
	"github.com/rigmas/microservices/product/handlers/product_grpc"
)

type Resolver struct {
	CustomerService customer_grpc.CustomerServiceClient
	ProductService  product_grpc.ProductServiceClient
	OrderService    order_grpc.OrderServiceClient
}
