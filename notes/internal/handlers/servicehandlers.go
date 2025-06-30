package handlers

import (
	"net/http"
	"sync/atomic"
)

var isShuttingDown atomic.Bool

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	if isShuttingDown.Load() {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("shutting down"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

}
