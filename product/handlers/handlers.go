package handlers

import (
	"context"

	"github.com/rigmas/microservices/product/handlers/product_grpc"
)

type Server struct {
	Products string

	product_grpc.UnimplementedProductServiceServer
}

func (s *Server) ProductList(_ context.Context,
	req *product_grpc.ProductListRequest) (*product_grpc.ProductListResponse, error) {

	var products []*product_grpc.Product

	product1 := &product_grpc.Product{
		Id:          "1",
		Title:       "Iphone",
		Description: "Iphone 13 Pro",
		Price:       18_000_000,
		Quantity:    20,
		CreatedAt:   "2022-08-15 17:01:36.399357+03",
	}

	product2 := &product_grpc.Product{
		Id:          "2",
		Title:       "Airpods",
		Description: "Airpods Pro 2",
		Price:       4_000_000,
		Quantity:    40,
		CreatedAt:   "2022-08-16 17:01:36.399357+03",
	}

	products = append(products, product1, product2)
	return &product_grpc.ProductListResponse{Products: products}, nil
}
