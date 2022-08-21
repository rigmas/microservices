package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/rigmas/microservices/customer/handlers/customer_grpc"
	"github.com/rigmas/microservices/gateway/graph/generated"
	"github.com/rigmas/microservices/gateway/graph/model"
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

	res := model.RegisterResponse{
		Token: resp.Token,
	}

	return &res, nil
}

// Ping is the resolver for the ping field.
func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	var products []*model.Product

	product1 := &model.Product{
		ID:          "1",
		Title:       "Iphone",
		Description: "Iphone 13 Pro",
		Price:       18_000_000,
		Quantity:    20,
		CreatedAt:   "2022-08-15 17:01:36.399357+03",
	}

	product2 := &model.Product{
		ID:          "2",
		Title:       "Airpods",
		Description: "Airpods Pro 2",
		Price:       4_000_000,
		Quantity:    40,
		CreatedAt:   "2022-08-16 17:01:36.399357+03",
	}

	products = append(products, product1, product2)

	return products, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
