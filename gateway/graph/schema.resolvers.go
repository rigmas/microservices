package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rigmas/microservices/customer/handlers/customer_grpc"
	"github.com/rigmas/microservices/gateway/graph/generated"
	"github.com/rigmas/microservices/gateway/graph/model"
	"github.com/rigmas/microservices/order/handlers/order_grpc"
	"github.com/rigmas/microservices/product/handlers/product_grpc"
)

// Health is the resolver for the health field.
func (r *mutationResolver) Health(ctx context.Context) (*model.HealthcheckResponse, error) {
	return &model.HealthcheckResponse{
		Message:     "ready",
		Healthiness: "healthy",
	}, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, username string, password string) (*model.RegisterResponse, error) {
	resp, err := r.CustomerService.Register(ctx, &customer_grpc.RegisterRequest{
		Username: username,
		Password: password,
	})

	if err != nil {
		return nil, err
	}

	res := &model.RegisterResponse{
		Token: resp.Token,
	}

	return res, nil
}

// Ping is the resolver for the ping field.
func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product

	resp, err := r.ProductService.ProductList(ctx, &product_grpc.ProductListRequest{})
	if err != nil {
		return nil, err
	}

	for _, product := range resp.Products {
		products = append(products, &model.Product{
			ID:          product.Id,
			Title:       product.Title,
			Description: product.Description,
			Price:       int(product.Price),
			Quantity:    int(product.Quantity),
			CreatedAt:   product.CreatedAt,
		})
	}

	return products, nil
}

// Orders is the resolver for the orders field.
func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var orders []*model.Order
	var products []*model.Product

	resp, err := r.OrderService.OrderList(ctx,
		&order_grpc.OrderListRequest{UserId: "1"},
	)
	if err != nil {
		return nil, err
	}

	for _, order := range resp.Orders {
		for _, product := range order.Products {
			products = append(products, &model.Product{
				ID:          product.Id,
				Title:       product.Title,
				Description: product.Description,
				Price:       int(product.Price),
				Quantity:    int(product.Quantity),
				CreatedAt:   product.CreatedAt,
			})
		}
		orders = append(orders, &model.Order{
			InvoiceNumber: order.InvoiceNumber,
			PaymentAmount: int(order.PaymentAmount),
			CreatedAt:     order.CreatedAt,
			OrderItems:    products,
		})
	}

	return orders, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
