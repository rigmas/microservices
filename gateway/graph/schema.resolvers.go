package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

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

// Ping is the resolver for the ping field.
func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
