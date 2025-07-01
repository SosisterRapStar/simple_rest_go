package httpserver

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/go-chi/chi"
	srg "github.com/sosisterrapstar/simple_rest_go"
)

var (
	isShuttingDown       atomic.Bool
	_shutdownPeriod      = 15 * time.Second
	_shutdownHardPeriod  = 3 * time.Second
	_readinessDrainDelay = 5 * time.Second
)

func readinessHandler(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug("Readiness handler check")
		if isShuttingDown.Load() {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("shutting down"))
			log.Debug("App is shutting down")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}
}

type Server struct {
	*http.Server
	log                   *slog.Logger
	cancelOngoingRequests context.CancelFunc
	// router                http.Handler
}

func New(cfg *srg.Config, router *chi.Mux, log *slog.Logger) *Server {
	address := net.JoinHostPort(cfg.Server.Host, cfg.Server.Port)
	baseCtx, cancel := context.WithCancel(context.Background())

	// special handlers
	router.Get("/healthz", readinessHandler(log))

	s := &http.Server{
		Addr:    address,
		Handler: router,
		BaseContext: func(_ net.Listener) context.Context {
			return baseCtx
		},
	}
	return &Server{
		Server:                s,
		log:                   log,
		cancelOngoingRequests: cancel,
	}
}

func (s *Server) Start() {
	go func() {
		s.log.Debug(fmt.Sprintf("Starting the server on %v ", s.Addr))
		if err := s.ListenAndServe(); err != nil {
			log.Fatal("Error starting the server")
		}
	}()
}

func (s *Server) Stop(ctx context.Context) {
	s.cancelOngoingRequests()

	isShuttingDown.Store(true)
	s.log.Debug("Received shutdown signal, shutting down.")

	time.Sleep(_readinessDrainDelay)
	s.log.Debug("Readiness check propagated, now waiting for ongoing requests to finish.")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), _shutdownPeriod)
	defer cancel()
	err := s.Shutdown(shutdownCtx)
	s.cancelOngoingRequests()
	if err != nil {
		log.Println("Failed to wait for ongoing requests to finish, waiting for forced cancellation.")
		time.Sleep(_shutdownHardPeriod)
	}
	s.log.Debug("Server shut down gracefully.")
}
