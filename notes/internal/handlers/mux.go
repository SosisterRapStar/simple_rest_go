package handlers

import "net/http"

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", mux.ServeHTTP)
	return mux
}
