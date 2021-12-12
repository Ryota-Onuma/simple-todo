package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"remote-schema/graph/generated"
	"remote-schema/graph/gqlmodel"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input gqlmodel.NewTask) (*gqlmodel.TaskCreateResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
