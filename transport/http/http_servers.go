package server

import (
	"net/http"
	"time"
)

func NewServer() http.Server {
	server := http.Server{
		Addr:           "0.0.0.0:8888",
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: (1 << 20),
	}
}
