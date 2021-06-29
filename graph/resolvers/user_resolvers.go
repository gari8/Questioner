package resolvers

import (
	"context"
	"errors"
	"faves4/graph/model"
	"faves4/infrastructure/auth"
	"faves4/infrastructure/lib/tools"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.Repository.UserRepository.NewUser(ctx, input)
}

func (r *mutationResolver) CreateSession(ctx context.Context, input model.LoginInput) (string, error) {
	s, err := r.Repository.NewSession(ctx, input)
	if err != nil {
		return "", err
	}

	if s == nil {
		return "", nil
	}

	return *s, nil
}

func (r *mutationResolver) EditUser(ctx context.Context, input model.EditUser) (*model.User, error) {
	return r.Repository.UpdateUser(ctx, input)
}

func (r *mutationResolver) EditPassword(ctx context.Context, id string, password string) (bool, error) {
	if _, err := r.Repository.UpdatePassword(ctx, id, password); err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) ([]*model.User, error) {
	l, o := 12, 0
	if limit != nil {
		l = *limit
	}
	if offset != nil {
		o = *offset
	}
	return r.Repository.FetchUsers(ctx, l, o)
}

func (r *queryResolver) FindUser(ctx context.Context, id string) (*model.User, error) {
	//ut, err := r.DB.User.Get(ctx, id)
	//if err != nil { return nil, err }
	//u := tools.CastUser(ut)
	//uq, err := ut.QueryQuestions().All(ctx)
	//if err != nil { return nil, err }
	//var qs []*model.Question
	//for _, qt := range uq {
	//	q, err := tools.CastQuestion(ctx, qt)
	//	if err != nil { continue }
	//	qs = append(qs, q)
	//}
	//u.Questions = qs
	//return u, nil
	return r.Repository.FetchUser(ctx, &id, nil)
}

func (r *queryResolver) ConfirmToken(ctx context.Context) (*model.User, error) {
	ut, ok := ctx.Value(auth.UserCtxKey).(string)
	if !ok { return nil, errors.New("not authenticate") }
	email, err := tools.DecodeToken(ut)
	if err != nil {
		return nil, err
	}
	return r.Repository.FetchUser(ctx, nil, &email)
}
