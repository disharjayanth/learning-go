package identity

import (
	"context"
	"net/http"
)

type userKey int

const key userKey = 1

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, key, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(key).(string)
	return user, ok
}

func extractUser(r *http.Request) (string, error) {
	userCookie, err := r.Cookie("user")
	if err != nil {
		return "", err
	}
	return userCookie.Value, nil
}

func Middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := extractUser(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := r.Context()
		ctx = ContextWithUser(ctx, user)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	}
}
