package handlers

import (
	"context"
	"time"

	"github.com/rigmas/microservices/order/handlers/order_grpc"
	"github.com/rigmas/microservices/product/handlers/product_grpc"
)

type Server struct {
	Orders string

	ProductClient product_grpc.ProductServiceClient

	order_grpc.UnimplementedOrderServiceServer
}

func (s *Server) OrderList(_ context.Context,
	req *order_grpc.OrderListRequest) (*order_grpc.OrderListResponse, error) {

	var orders []*order_grpc.Order
	var products []*order_grpc.Product

	reqProduct := &product_grpc.ProductListRequest{}

	resp, err := s.ProductClient.ProductList(context.Background(), reqProduct)
	if err != nil {
		return nil, err
	}

	var totalAmount int
	for _, product := range resp.Products {

		products = append(products, &order_grpc.Product{
			Id:          product.Id,
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
			Quantity:    product.Quantity,
			CreatedAt:   product.CreatedAt,
		})

		totalAmount += (int(product.Price) * int(product.Quantity))
	}

	order := &order_grpc.Order{
		InvoiceNumber: "TEST123XXX",
		PaymentAmount: int32(totalAmount),
		CreatedAt:     time.Now().String(),
		Products:      products,
	}

	orders = append(orders, order)

	return &order_grpc.OrderListResponse{Orders: orders}, nil
}
