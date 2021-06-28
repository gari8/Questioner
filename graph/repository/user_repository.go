package repository

import (
	"context"
	"faves4/ent"
	"faves4/ent/user"
	"faves4/graph/model"
	"faves4/infrastructure/lib/tools"
)

type (
	userRepository struct {
		*ent.Client
	}
	UserRepository interface {
		FetchUser(ctx context.Context, id string) (*model.User, error)
	}
)

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client}
}

func (u *userRepository) FetchUser(ctx context.Context, id string) (*model.User, error) {
	ut, err := u.Client.User.Query().
		Where(user.IDEQ(id)).WithQuestions(func(q *ent.QuestionQuery) {
			q.WithAnswers()
			q.WithChoices()
		}).Only(ctx)
	if err != nil { return nil, err }
	us := tools.CastUser(ut)
	return us, nil
}