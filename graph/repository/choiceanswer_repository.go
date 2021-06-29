package repository

import (
	"context"
	"faves4/ent"
)

type (
	choiceAnswerRepository struct {
		*ent.Client
	}
	ChoiceAnswerRepository interface {
		FetchChoiceAnswer(ctx context.Context, id string) error
	}
)

func NewChoiceAnswerRepository(client *ent.Client) ChoiceAnswerRepository {
	return &choiceAnswerRepository{client}
}

func (q *choiceAnswerRepository) FetchChoiceAnswer(ctx context.Context, id string) error {
	return nil
}