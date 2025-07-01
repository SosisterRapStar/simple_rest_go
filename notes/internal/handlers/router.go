package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

func NewMux(log *slog.Logger) http.Handler {
	mux := chi.NewMux()
	return mux
}
