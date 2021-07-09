package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/andrewtanjaya21/test_go/graph/generated"
	"github.com/andrewtanjaya21/test_go/graph/model"
)

func (r *mutationResolver) CreateFood(ctx context.Context, input model.NewFood) (*model.Food, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateFood(ctx context.Context, id string, input model.NewFood) (*model.Food, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteFood(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Foods(ctx context.Context) ([]*model.Food, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
