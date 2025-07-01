package middlewares

import (
	"context"
	"log/slog"
	"net/http"
)

type UserId struct{}

var UserIdKey = UserId{}

func UserIdFromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(UserIdKey).(string)
	if !ok {
		return "", false
	}
	return val, true
}

func AuthMiddleware(logger slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		})
	}
}
