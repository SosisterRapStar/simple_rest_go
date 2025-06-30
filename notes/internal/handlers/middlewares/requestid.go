package middlewares

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

type CorrelationId struct{}

var CorId CorrelationId = CorrelationId{}

func CorrelationIdFromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(CorId).(string)
	if !ok {
		return "", false
	}
	return id, ok
}

// RequestIdMIddleware - middleware for getting requestId from the request,
// should always be the first middleware in middleware chain,
// if there is no any request id from request, logger will log it like warn event
func RequestIdMIddleware(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := r.Header.Get("X-Correlation-ID")
			if req == "" {
				logger.Warn("request without correlation id, will be created a new id")
				req = uuid.New().String()
			}
			r = r.WithContext(context.WithValue(context.Background(), CorId, req))
			next.ServeHTTP(w, r)
		})
	}
}
