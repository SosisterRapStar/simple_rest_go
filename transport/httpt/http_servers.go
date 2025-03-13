package httpt

import (
	"context"
	"errors"
	"first-proj/appconfig"
	"fmt"
	"log"
	"net/http"
	"time"
)

var logger = appconfig.GetLogger()

type Server struct {
	server *http.Server
	Router *http.ServeMux
}

func loggingMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Info(fmt.Sprintf("Server started to server request %s %s", r.Method, r.URL.Path))
		next.ServeHTTP(w, r)
		logger.Info(fmt.Sprintf("Request serverd for %v", time.Since(start)))
	})
}

func NewServer(address string, handlers HttpApi) Server {
	logger.Debug("New server is activating")

	router := http.NewServeMux()
	wrapedRouter := loggingMidleware(router)
	server := http.Server{
		Addr:           address,
		Handler:        wrapedRouter,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: (1 << 20),
	}

	router.HandleFunc("GET /api/v1/notes/{id}", handlers.GetNoteById)
	router.HandleFunc("POST /api/v1/notes", handlers.CreateNote)
	router.HandleFunc("PATCH /api/v1/notes/{id}", handlers.UpdateNote)
	router.HandleFunc("DELETE /api/v1/notes/{id}", handlers.DeleteNote)
	router.HandleFunc("GET /api/v1/notes/", handlers.FindNotes)

	return Server{
		server: &server,
		Router: router,
	}
}

func (s *Server) Start() {

	go func() {
		logger.Info("Starting the server")
		if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
		logger.Info("Stopped serving new connections.")
	}()
}

func (s *Server) Stop(ctx context.Context) {
	logger.Debug("Shutting down the server")
	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatal("Error occured shutting the server down")
	}
	logger.Info("Server was stopped")
}
