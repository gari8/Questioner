package resolvers

import (
	"faves4/graph/generated"
	"faves4/interactor"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{Repository interactor.Repository}
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func NewResolver(repository interactor.Repository) *Resolver {
	return &Resolver{
		Repository: repository,
	}
}

func (r *Resolver) Directive() generated.DirectiveResolver {
	panic("implement me")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

