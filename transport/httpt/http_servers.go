package httpt

import (
	"context"
	"first-proj/appconfig"
	"fmt"
	"net/http"
	"os"
	"time"
)

var logger = appconfig.GetLogger()

type Server struct {
	server *http.Server
	Router *http.ServeMux
}

func NewServer(address string) Server {
	logger.Debug("New server is activating")

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

func (s *Server) Start(done <-chan struct{}) {
	logger.Debug("Starting server")

	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			logger.Error("Error occured during the server starting", err)
			os.Exit(1)
		}
	}()
	fmt.Println("Server is up")
	<-done
	logger.Debug("Server is shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown error:", err)
	}
	fmt.Println("Server stopped")

}

func (s *Server) Stop() {
	logger.Debug("Stopping the server")
	err := s.server.Close()
	if err != nil {
		fmt.Println("Error occured stopping the server")
		os.Exit(1)
	}
	logger.Debug("Server is stopped")

}
