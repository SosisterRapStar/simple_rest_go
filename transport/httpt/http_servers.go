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
