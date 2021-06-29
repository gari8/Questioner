package resolvers

import (
	"context"
	"faves4/graph/model"
	"fmt"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.NewQuestion) (*model.Question, error) {
	return r.Repository.NewQuestion(ctx, input)
}

func (r *mutationResolver) EditQuestion(ctx context.Context, input model.EditQuestion) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Questions(ctx context.Context, limit *int, offset *int) ([]*model.Question, error) {
	l, o := 12, 0
	if limit != nil {
		l = *limit
	}
	if offset != nil {
		o = *offset
	}
	return r.Repository.FetchQuestions(ctx, l, o)
}

func (r *queryResolver) FindQuestion(ctx context.Context, id string, userId *string) (*model.Question, error) {
	return r.Repository.FetchQuestion(ctx, id, userId)
}
