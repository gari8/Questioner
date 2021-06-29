package repository

import (
	"context"
	"faves4/ent"
	"faves4/graph/model"
)

type (
	choiceRepository struct {
		*ent.Client
	}
	ChoiceRepository interface {
		FetchChoice(ctx context.Context, id string) (*model.Choice, error)
	}
)

func NewChoiceRepository(client *ent.Client) ChoiceRepository {
	return &choiceRepository{client}
}

func (q *choiceRepository) FetchChoice(ctx context.Context, id string) (*model.Choice, error) {
	return nil, nil
}