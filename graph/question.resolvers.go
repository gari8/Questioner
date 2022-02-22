package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gari8/Questioner/domain"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, newQuestionInput *domain.NewQuestionInput) (*domain.Question, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditQuestion(ctx context.Context, editQuestionInput *domain.EditQuestionInput) (*domain.Question, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Questions(ctx context.Context, keyword *string, first *int, offset *int) ([]*domain.Question, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) QuestionByID(ctx context.Context, id string) (*domain.Question, error) {
	panic(fmt.Errorf("not implemented"))
}
