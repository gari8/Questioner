package resolvers

import (
	"context"
	"faves4/graph/model"
	"faves4/infrastructure/lib/tools"
	"fmt"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, input model.NewQuestion) (*model.Question, error) {
	uu, err := tools.CreateId()
	if err != nil { return nil, err }
	qs, err := r.DB.Question.Create().SetID(*uu).SetOwnerID(input.UserID).SetTitle(input.Title).SetContent(input.Content).SetTextAfterAnswered(*input.TextAfterAnswered).SetAnswerType(input.AnswerType.String()).SetEnabled(input.Enabled).Save(ctx)
	if err != nil { return nil, err }
	qt, err := tools.CastQuestion(ctx, qs)
	if err != nil { return nil, err }
	if input.AnswerType == model.AnswerTypeSelect {
		var choices []*model.Choice
		for _, cc := range input.Choices {
			ct, err := r.DB.Choice.Create().SetParentID(qs.ID).SetContent(cc.Content).Save(ctx)
			if err != nil { continue }
			c, err := tools.CastChoice(ctx, ct)
			if err != nil { continue }
			choices = append(choices, c)
		}
		qt.Choices = choices
	}
	return qt, nil
}

func (r *mutationResolver) EditQuestion(ctx context.Context, input model.EditQuestion) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Questions(ctx context.Context, limit *int, offset *int) ([]*model.Question, error) {
	qAll, err := r.DB.Question.Query().Offset(*offset).Limit(*limit).All(ctx)
	if err != nil { return nil, err }
	var qs []*model.Question
	for _, qt := range qAll {
		q, err := tools.CastQuestion(ctx, qt)
		if err != nil { continue }
		qs = append(qs, q)
	}
	return qs, nil
}

func (r *queryResolver) FindQuestion(ctx context.Context, id string, userId *string) (*model.Question, error) {
	qt, err := r.DB.Question.Get(ctx, id)
	if err != nil { return nil, err }
	q, err := tools.CastQuestion(ctx, qt)
	if err != nil { return nil, err }
	aAll, err := qt.QueryAnswers().All(ctx)
	if err != nil { return nil, err }
	cAll, err := qt.QueryChoices().All(ctx)
	if err != nil { return nil, err }
	var answers []*model.Answer
	var answerers []*model.User
	var choices []*model.Choice
	for _, ans := range aAll {
		a, err := tools.CastAnswer(ctx, ans)
		if err != nil { continue }
		answers = append(answers, a)
		answerers = append(answerers, a.User)
	}
	ans := false
	for _, choice := range cAll {
		tmpId, err := choice.QueryChoiceanswers().QueryOwner().FirstID(ctx)
		if userId != nil && err == nil {
			if tmpId == *userId { ans = true }
		}
		c, err := tools.CastChoice(ctx, choice)
		if err != nil { continue }
		count, err := choice.QueryChoiceanswers().Count(ctx)
		if err != nil {
			c.Value = 0
		} else {
			c.Value = count
		}
		choices = append(choices, c)
	}
	q.Answered = ans
	q.Answerers = answerers
	q.Answers = answers
	q.Choices = choices
	if (q.AnswerType != string(model.AnswerTypeSelect)) && (userId != nil) {
		q.Answered = tools.ContainsUserIdInAnswers(answers, *userId)
	}
	return q, nil
}
