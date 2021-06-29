package repository

import (
	"context"
	"faves4/ent"
	"faves4/graph/model"
)

type (
	answerRepository struct {
		*ent.Client
	}
	AnswerRepository interface {
		FetchAnswer(ctx context.Context, id string) (*model.Answer, error)
		NewAnswer(ctx context.Context, input model.NewAnswer) (bool, error)
	}
)

func NewAnswerRepository(client *ent.Client) AnswerRepository {
	return &answerRepository{client}
}

func (a *answerRepository) FetchAnswer(ctx context.Context, id string) (*model.Answer, error) {
	return nil, nil
}

func (a *answerRepository) NewAnswer(ctx context.Context, input model.NewAnswer) (bool, error) {
	if input.AnswerType == model.AnswerTypeSelect {
		if input.ChoiceID == nil { return false, nil }
		_, err := a.ChoiceAnswer.Create().SetOwnerID(input.UserID).SetParentID(*input.ChoiceID).Save(ctx)
		if err != nil { return false, err }
		return true, nil
	}
	_, err := a.Answer.Create().SetOwnerID(input.UserID).SetParentID(input.QuestionID).SetAnswerType(string(input.AnswerType)).SetContent(*input.Value).Save(ctx)
	if err != nil { return false, err }
	return true, nil
}
