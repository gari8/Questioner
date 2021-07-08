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

func (r *mutationResolver) EditPassword(ctx context.Context, id string, newPassword string, currentPassword string) (bool, error) {
	if ok := r.Repository.ConfirmPassword(ctx, id, currentPassword); !ok {
		return false, errors.New("the key pair didn't match")
	}
	if _, err := r.Repository.UpdatePassword(ctx, id, newPassword); err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Users(ctx context.Context, limit *int, offset *int) (*model.UsersOutput, error) {
	l, o := 12, 0
	if limit != nil {
		l = *limit
	}
	if offset != nil {
		o = *offset
	}
	us, err := r.Repository.FetchUsers(ctx, l, o)
	if err != nil { return nil, err }
	return &model.UsersOutput{
		Users: us,
		Length: r.Repository.GetAllUsersCount(ctx),
	}, nil
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
