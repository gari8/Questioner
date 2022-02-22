package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gari8/Questioner/domain"
)

func (r *mutationResolver) EditUser(ctx context.Context, editUserInput *domain.EditUserInput) (*domain.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, keyword *string, first *int, offset *int) ([]*domain.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserByID(ctx context.Context, id string) (*domain.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic(fmt.Errorf("not implemented"))
}
