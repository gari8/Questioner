package auth

import (
	"context"
	"github.com/gari8/Questioner/domain"
	"net/http"
)

const UserCtxKey = "userToken"

func Middleware(ctx context.Context, repository interactor.Repository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if val, ok := r.Header["Authorization"]; ok {
				if len(val) != 0 {
					c := context.WithValue(r.Context(), UserCtxKey, val[0])
					r = r.WithContext(c)
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *domain.User {
	//raw, _ := ctx.Value(userCtxKey).(*model.User)
	raw := &domain.User{}
	return raw
}
