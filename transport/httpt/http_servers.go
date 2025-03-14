package httpt

import (
	"context"
	"errors"
	"first-proj/appconfig"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var logger = appconfig.GetLogger()

type Server struct {
	server *http.Server
	Router *http.ServeMux
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func reportPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, "Internal error occured", http.StatusInternalServerError)
				logger.Error("Panic occured:", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// logs time of response, creates prom record
func loggingMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		activeRequestsGauge.Inc()
		start := time.Now()
		logger.Info(fmt.Sprintf("Server started to server request %s %s", r.Method, r.URL.Path))
		customRecorder := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		duration := time.Since(start)

		next.ServeHTTP(customRecorder, r)
		activeRequestsGauge.Dec()
		method := r.Method
		path := r.URL.Path

		status := strconv.Itoa(customRecorder.status)
		latencyHistogram.With(prometheus.Labels{
			"method": method, "path": path, "status": status,
		}).Observe(duration.Seconds())
		httpRequestCounter.WithLabelValues(status, path, method).Inc()
	})
}

func NewServer(config *appconfig.Config, handlers HttpApi) Server {

	router := http.NewServeMux()
	wrapedRouter := loggingMidleware(router)
	wrapedRouter = reportPanic(wrapedRouter)

	server := http.Server{
		Addr:           config.Address,
		Handler:        wrapedRouter,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: (1 << 20),
	}

	newServer := Server{
		server: &server,
		Router: router,
	}
	newServer.RegisterRoutes(handlers)
	return newServer
}

func (s *Server) RegisterRoutes(handlers HttpApi) {
	s.Router.HandleFunc("GET /api/v1/notes/{id}", handlers.GetNoteById)
	s.Router.HandleFunc("POST /api/v1/notes", handlers.CreateNote)
	s.Router.HandleFunc("PATCH /api/v1/notes/{id}", handlers.UpdateNote)
	s.Router.HandleFunc("DELETE /api/v1/notes/{id}", handlers.DeleteNote)
	s.Router.HandleFunc("GET /api/v1/notes/", handlers.FindNotes)
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

// Prometheus metrics stuff
var (
	httpRequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of http requests received",
		}, []string{
			"status",
			"path",
			"method"})
	activeRequestsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_active_requests",
			Help: "Number of active connections to the service",
		},
	)
	latencyHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "http_request_duration_seconds",
		Help: "Duration of HTTP requests",
	}, []string{"status", "path", "method"})
)

func ListenAndServeProm(config *appconfig.Config) {
	mux := http.NewServeMux()
	reg := prometheus.NewRegistry()

	// metrics registration
	reg.MustRegister(httpRequestCounter)
	reg.MustRegister(activeRequestsGauge)
	reg.MustRegister(latencyHistogram)

	handler := promhttp.HandlerFor(
		reg,
		promhttp.HandlerOpts{},
	)
	mux.Handle("/metrics", handler)

	// Получаем порт из переменной окружения (или используем значение по умолчанию)
	port := config.MetricsPort

	log.Printf("Starting Prometheus metrics server on %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
