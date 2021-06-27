package resolvers

import (
	"context"
	"faves4/graph/model"
	"fmt"
)

func (r *mutationResolver) CreateAnswer(ctx context.Context, input model.NewAnswer) (bool, error) {
	if input.AnswerType == model.AnswerTypeSelect {
		if input.ChoiceID == nil { return false, nil }
		_, err := r.DB.ChoiceAnswer.Create().SetOwnerID(input.UserID).SetParentID(*input.ChoiceID).Save(ctx)
		if err != nil { return false, err }
		return true, nil
	}
	_, err := r.DB.Answer.Create().SetOwnerID(input.UserID).SetParentID(input.QuestionID).SetAnswerType(string(input.AnswerType)).SetContent(*input.Value).Save(ctx)
	if err != nil { return false, err }
	return true, nil
}

func (r *queryResolver) FindAnswer(ctx context.Context, id string) (*model.Answer, error) {
	panic(fmt.Errorf("not implemented"))
}
