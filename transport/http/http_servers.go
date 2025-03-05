package http

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	server *http.Server
	Router *http.ServeMux
}

func NewServer(address string) Server {
	router := http.NewServeMux()
	server := http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: (1 << 20),
	}
	return Server{
		server: &server,
		Router: router,
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
