package resolvers

import (
	"context"
	"errors"
	"faves4/ent/user"
	"faves4/graph/model"
	"faves4/infrastructure/auth"
	"faves4/infrastructure/lib/tools"
	"fmt"
	"time"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	uu, err := tools.CreateId()
	if err != nil { return nil, err }
	hash, err := tools.DigestWord(input.Password)
	if err != nil { return nil, err }
	u, err := r.DB.User.Create().SetID(*uu).SetUsername(input.Username).SetEmail(input.Email).SetPassword(hash).SetCratedAt(time.Now()).Save(ctx)
	if err != nil || u == nil { return nil, err }
	ca := u.CratedAt.String()
	ua := u.UpdatedAt.String()
	str := ""
	return &model.User{
		ID: u.ID,
		Username: u.Username,
		Email: &u.Email,
		Password: tools.MuskWord(input.Password),
		Icon: &str,
		Description: &str,
		CreatedAt: &ca,
		UpdatedAt: &ua,
	}, nil
}

func (r *mutationResolver) CreateSession(ctx context.Context, input *model.LoginInput) (string, error) {
	u, err := r.DB.User.Query().Where(user.EmailEQ(input.Email)).Only(ctx)
	if err != nil { return "", err }
	if !tools.CompareHashAndPlane(u.Password, input.Password) { return "", err }
	t, err := tools.CreateToken(u.Email)
	if err != nil { return t, err }
	return t, nil
}

func (r *mutationResolver) RemoveSession(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EditUser(ctx context.Context, input model.EditUser) (*model.User, error) {
	u, err := r.DB.User.Query().Where(user.EmailEQ(input.Email)).Only(ctx)
	if err != nil { return nil, err }
	u, err = u.Update().SetEmail(input.Email).SetUsername(input.Username).SetDescription(input.Description).SetUpdatedAt(time.Now()).Save(ctx)
	if err != nil { return nil, err }
	return tools.CastUser(u), nil
}

func (r *mutationResolver) EditPassword(ctx context.Context, id string, password string) (bool, error) {
	u, err := r.DB.User.Get(ctx, id)
	if err != nil { return false, nil }
	hash, err := tools.DigestWord(password)
	if err != nil { return false, nil }
	_, err = u.Update().SetPassword(hash).Save(ctx)
	if err != nil { return false, nil }
	return true, nil
}

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*model.User, error) {
	ul, err := r.DB.User.Query().Limit(*limit).Offset(*offset).All(ctx)
	if err != nil { return nil, err }
	var us []*model.User
	for _, u := range ul {
		us = append(us, tools.CastUser(u))
	}
	return us, nil
}

func (r *queryResolver) FindUser(ctx context.Context, id string) (*model.User, error) {
	ut, err := r.DB.User.Get(ctx, id)
	if err != nil { return nil, err }
	u := tools.CastUser(ut)
	uq, err := ut.QueryQuestions().All(ctx)
	if err != nil { return nil, err }
	var qs []*model.Question
	for _, qt := range uq {
		q, err := tools.CastQuestion(ctx, qt)
		if err != nil { continue }
		qs = append(qs, q)
	}
	u.Questions = qs
	return u, nil
}

func (r *queryResolver) ConfirmToken(ctx context.Context) (*model.User, error) {
	ut, ok := ctx.Value(auth.UserCtxKey).(string)
	if !ok { return nil, errors.New("not authenticate") }
	email, err := tools.DecodeToken(ut)
	if err != nil {
		return nil, err
	}
	u, err := r.DB.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	cu := tools.CastUser(u)
	return cu, nil
}
