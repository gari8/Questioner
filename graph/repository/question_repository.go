package repository

import (
	"context"
	"faves4/ent"
	"faves4/ent/choiceanswer"
	"faves4/ent/user"
	"faves4/graph/model"
	"faves4/infrastructure/lib/tools"
)

type (
	questionRepository struct {
		*ent.Client
	}
	QuestionRepository interface {
		FetchQuestion(ctx context.Context, id string, userId *string) (*model.Question, error)
		FetchQuestions(ctx context.Context, limit int, offset int) ([]*model.Question, error)
		NewQuestion(ctx context.Context, input model.NewQuestion) (*model.Question, error)
	}
)

func NewQuestionRepository(client *ent.Client) QuestionRepository {
	return &questionRepository{client}
}

func (q *questionRepository) FetchQuestion(ctx context.Context, id string, userId *string) (*model.Question, error) {
	qt, err := q.Question.Get(ctx, id)
	if err != nil { return nil, err }
	qs, err := tools.CastQuestion(ctx, qt)
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
	cnt := 0
	if userId == nil { ans = true }
	for _, choice := range cAll {
		if userId != nil {
			tmpId, _ := choice.QueryChoiceanswers().
				Where(choiceanswer.HasOwnerWith(user.ID(*userId))).Count(ctx)
			if tmpId > 0 { ans = true }
		}
		c, err := tools.CastChoice(ctx, choice)
		if err != nil { continue }
		count, err := choice.QueryChoiceanswers().Count(ctx)
		if err != nil {
			c.Value = 0
		} else {
			c.Value = count
		}
		cnt += count
		choices = append(choices, c)
		cau, _ := choice.QueryChoiceanswers().QueryOwner().Only(ctx)
		if cau == nil { continue }
		answerers = append(answerers, tools.CastUser(cau))
	}
	qs.Answered = ans
	qs.Answerers = answerers
	qs.Answers = answers
	qs.Choices = choices
	qs.AnswerCount = len(answers) + cnt
	if qs.AnswerType != string(model.AnswerTypeSelect) {
		if userId != nil {
			qs.Answered = tools.ContainsUserIdInAnswers(answers, *userId)
		}
	}
	return qs, nil
}

func (q *questionRepository) NewQuestion(ctx context.Context, input model.NewQuestion) (*model.Question, error) {
	uu, err := tools.CreateId()
	if err != nil { return nil, err }
	qs, err := q.Question.Create().SetID(*uu).
		SetOwnerID(input.UserID).SetTitle(input.Title).
		SetContent(input.Content).SetTextAfterAnswered(*input.TextAfterAnswered).
		SetAnswerType(input.AnswerType.String()).SetEnabled(input.Enabled).
		Save(ctx)
	if err != nil { return nil, err }
	qt, err := tools.CastQuestion(ctx, qs)
	if err != nil { return nil, err }
	if input.AnswerType == model.AnswerTypeSelect {
		var choices []*model.Choice
		for _, cc := range input.Choices {
			ct, err := q.Choice.Create().SetParentID(qs.ID).SetContent(cc.Content).Save(ctx)
			if err != nil { continue }
			c, err := tools.CastChoice(ctx, ct)
			if err != nil { continue }
			choices = append(choices, c)
		}
		qt.Choices = choices
	}
	return qt, nil
}

func (q *questionRepository) FetchQuestions(ctx context.Context, limit int, offset int) ([]*model.Question, error) {
	qAll, err := q.Question.Query().Offset(offset).Limit(limit).All(ctx)
	if err != nil { return nil, err }
	var qs []*model.Question
	for _, qt := range qAll {
		q, err := tools.CastQuestion(ctx, qt)
		if err != nil { continue }
		qs = append(qs, q)
	}
	return qs, nil
}
