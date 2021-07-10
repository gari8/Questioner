package resolvers

import (
	"context"
	"faves4/graph/model"
)

func (r *queryResolver) Search(ctx context.Context, keyword string) (*model.SearchOutput, error) {
	var users []*model.User
	var questions []*model.Question
	return &model.SearchOutput{
		Users: users,
		Questions: questions,
	}, nil
}