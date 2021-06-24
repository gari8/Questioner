package resolvers

import (
	"context"
	"fmt"
	"server/graph/model"
)

func (r *mutationResolver) CreateAnswer(ctx context.Context, input model.NewAnswer) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FindAnswer(ctx context.Context, id string) (*model.Answer, error) {
	panic(fmt.Errorf("not implemented"))
}
