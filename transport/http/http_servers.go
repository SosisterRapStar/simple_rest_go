package server

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	server *http.Server
	router *http.ServeMux
	api    *HttpApi
}

func NewServer() Server {
	router := http.NewServeMux()
	server := http.Server{
		Addr:           "0.0.0.0:8888",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: (1 << 20),
	}
	return Server{
		server: &server,
		router: router,
	}
}

func (s *Server) Start() {
	err := s.server.ListenAndServe()
	if err != nil {
		fmt.Println("Error occured during the server starting")
		os.Exit(1)
	}
	fmt.Println("Server is up")
}

func (s *Server) Stop() {
	err := s.server.Close()
	if err != nil {
		fmt.Println("Error occured stopping the server")
		os.Exit(1)
	}
	fmt.Println("Server is stopped")
}

func (s *Server) RegisterRouters() {
	s.router.HandleFunc("/", s.api.CreateNote)
	// ...
}
