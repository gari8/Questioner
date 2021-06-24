package resolvers

import (
	"context"
	"faves4/graph/model"
	"fmt"
)

func (r *mutationResolver) CreateAnswer(ctx context.Context, input model.NewAnswer) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FindAnswer(ctx context.Context, id string) (*model.Answer, error) {
	panic(fmt.Errorf("not implemented"))
}
