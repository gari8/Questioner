package repository

import (
	"context"
	"faves4/ent"
	"faves4/ent/question"
	"faves4/ent/user"
	"faves4/graph/model"
	"faves4/infrastructure/lib/tools"
	"time"
)

type (
	userRepository struct {
		*ent.Client
	}
	UserRepository interface {
		FetchUser(ctx context.Context, id *string, email *string) (*model.User, error)
		NewUser(ctx context.Context, input model.NewUser) (*model.User, error)
		NewSession(ctx context.Context, input model.LoginInput) (*string, error)
		UpdateUser(ctx context.Context, input model.EditUser) (*model.User, error)
		UpdatePassword(ctx context.Context, id string, password string) (*model.User, error)
		FetchUsers(ctx context.Context, limit int, offset int) ([]*model.User, error)
		ConfirmPassword(ctx context.Context, id string, password string) bool
		GetAllUsersCount(ctx context.Context) int
	}
)

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client}
}

func (u *userRepository) NewUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	uu, err := tools.CreateId()
	if err != nil { return nil, err }
	hash, err := tools.DigestWord(input.Password)
	if err != nil { return nil, err }
	us, err := u.User.Create().SetID(*uu).SetUsername(input.Username).SetEmail(input.Email).SetPassword(hash).Save(ctx)
	if err != nil || u == nil { return nil, err }
	ca := us.CratedAt.String()
	ua := us.UpdatedAt.String()
	return &model.User{
		ID: us.ID,
		Username: us.Username,
		Email: &us.Email,
		Password: tools.MuskWord(input.Password),
		Icon: &us.Icon,
		Description: &us.Description,
		CreatedAt: &ca,
		UpdatedAt: &ua,
	}, nil
}

func (u *userRepository) FetchUser(ctx context.Context, id *string, email *string) (*model.User, error) {
	var ut *ent.User
	var uqs []*model.Question
	var err error
	if id != nil {
		ut, err = u.User.Query().
			Where(user.IDEQ(*id)).WithQuestions(func(q *ent.QuestionQuery) {
			q.WithOwner()
			q.WithAnswers()
			q.WithChoices()
			q.Order(ent.Desc(question.FieldCratedAt))
		}).WithAnswers().WithChoiceanswers().Only(ctx)
	}
	if email != nil {
		ut, err = u.User.Query().
			Where(user.EmailEQ(*email)).WithQuestions(func(q *ent.QuestionQuery) {
			q.WithOwner()
			q.WithAnswers()
			q.WithChoices()
		}).WithAnswers().WithChoiceanswers().Only(ctx)
	}
	if ut != nil {
		for _, qType := range ut.Edges.Questions {
			qs, _ := tools.CastQuestion(ctx, qType)
			if qs != nil {
				qs.User = tools.CastUser(qType.Edges.Owner)
				uqs = append(uqs, qs)
			}
		}
	}
	if err != nil { return nil, err }
	us := tools.CastUser(ut)
	al := len(ut.Edges.Answers) + len(ut.Edges.Choiceanswers)
	ql := len(ut.Edges.Questions)
	us.AnswerCount = &al
	us.QuestionCount = &ql
	us.Questions = uqs
	return us, nil
}

func (u *userRepository) UpdateUser(ctx context.Context, input model.EditUser) (*model.User, error) {
	us, err := u.User.UpdateOneID(input.ID).
		SetEmail(input.Email).SetUsername(input.Username).
		SetDescription(input.Description).SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil { return nil, err }
	return tools.CastUser(us), nil
}

func (u *userRepository) UpdatePassword(ctx context.Context, id string, password string) (*model.User, error) {
	us, err := u.User.Get(ctx, id)
	if err != nil { return nil, err }
	hash, err := tools.DigestWord(password)
	if err != nil { return nil, err }
	us, err = us.Update().SetPassword(hash).Save(ctx)
	if err != nil { return nil, err }
	return tools.CastUser(us), nil
}

func (u *userRepository) NewSession(ctx context.Context, input model.LoginInput) (*string, error) {
	us, err := u.User.Query().Where(user.EmailEQ(input.Email)).Only(ctx)
	if err != nil { return nil, err }
	if !tools.CompareHashAndPlane(us.Password, input.Password) { return nil, err }
	t, err := tools.CreateToken(us.Email)
	if err != nil { return nil, err }
	return &t, nil
}

func (u *userRepository) FetchUsers(ctx context.Context, limit int, offset int) ([]*model.User, error) {
	ul, err := u.User.Query().WithAnswers().WithQuestions().WithChoiceanswers().Limit(limit).Offset(offset).All(ctx)
	if err != nil { return nil, err }
	var us []*model.User
	for _, ut := range ul {
		um := tools.CastUser(ut)
		al := len(ut.Edges.Answers) + len(ut.Edges.Choiceanswers)
		ql := len(ut.Edges.Questions)
		um.AnswerCount = &al
		um.QuestionCount = &ql
		us = append(us, um)
	}
	return us, nil
}

func (u *userRepository) ConfirmPassword(ctx context.Context, id string, password string) bool {
	us, err := u.User.Get(ctx, id)
	if err != nil {
		return false
	}
	if ok := tools.CompareHashAndPlane(us.Password, password); ok { return true }
	return false
}

func (u *userRepository) GetAllUsersCount(ctx context.Context) int {
	return u.User.Query().CountX(ctx)
}