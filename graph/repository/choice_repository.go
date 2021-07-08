package repository

import (
	"context"
	"faves4/ent"
	"faves4/ent/choice"
	"faves4/graph/model"
	"faves4/infrastructure/lib/tools"
)

type (
	choiceRepository struct {
		*ent.Client
	}
	ChoiceRepository interface {
		FetchChoice(ctx context.Context, id int, userId *string) (*model.Choice, error)
	}
)

func NewChoiceRepository(client *ent.Client) ChoiceRepository {
	return &choiceRepository{client}
}

func (q *choiceRepository) FetchChoice(ctx context.Context, id int, userId *string) (*model.Choice, error) {
	c, err := q.Choice.Query().Where(choice.IDEQ(id)).
		WithChoiceanswers(func(query *ent.ChoiceAnswerQuery) {
			query.WithOwner()
		}).
		WithParent().Only(ctx)
	if err != nil {
		return nil, err
	}
	ct, err := tools.CastQuestion(ctx, c.Edges.Parent)
	if err != nil {
		return nil, err
	}

	answered := false

	for _, v := range c.Edges.Choiceanswers {
		if v.Edges.Owner.ID == *userId { answered = true }
	}

	return &model.Choice{
		ID: id,
		Content: c.Content,
		Value: len(c.Edges.Choiceanswers),
		Question: ct,
		Answered: answered,
	}, nil
}