package auth

import (
	"context"
	"faves4/ent"
	"faves4/graph/model"
	"net/http"
)

const UserCtxKey = "userToken"

func Middleware(ctx context.Context, client *ent.Client) func(http.Handler) http.Handler {
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

func ForContext(ctx context.Context) *model.User {
	//raw, _ := ctx.Value(userCtxKey).(*model.User)
	raw := &model.User{}
	return raw
}