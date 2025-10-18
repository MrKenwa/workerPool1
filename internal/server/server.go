package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"go.uber.org/zap"
	workerPoolHandler "workerPool1/internal/api/worker_pool"
	"workerPool1/internal/config"
	"workerPool1/internal/logger"
	workerPoolService "workerPool1/internal/service/worker_pool"
)

type Server struct {
	srv *http.Server
	mux *http.ServeMux
	cfg *config.Config
	log *logger.Logger
}

func New(cfg *config.Config, log *logger.Logger) *Server {
	mux := &http.ServeMux{}
	return &Server{
		srv: &http.Server{
			Addr:    cfg.ServerConfig.Host,
			Handler: mux,
		},
		mux: mux,
		cfg: cfg,
		log: log,
	}
}

func (s *Server) Start() error {
	// Оборачиваем mux в middleware
	s.srv.Handler = s.loggingMiddleware(s.mux)

	workerPool, err := workerPoolService.New(&s.cfg.QueueConfig, s.log)
	if err != nil {
		s.log.Errorf("error while creating worker pool: %v", err)
		return err
	}
	go workerPool.Start()

	wpHandler := workerPoolHandler.New(workerPool, s.log)
	workerPoolHandler.MapRoutes(s.mux, wpHandler)

	s.mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	go func() {
		s.log.Infof("Server started on %s", s.cfg.ServerConfig.Host)
		if err := s.srv.ListenAndServe(); err != nil {
			s.log.Errorf("Listen error: %v", err)
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	s.log.Infof("Shutting down...")

	if err = workerPool.Close(); err != nil {
		s.log.Errorf("error while closing worker pool: %v", err)
	}
	log.Print("Worker pool is closed")

	if err = s.srv.Shutdown(context.Background()); err != nil {
		s.log.Errorf("server shutdown error: %v", err)
		return err
	}
	s.log.Infof("Server is stopped")

	return nil
}

func (s *Server) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// health-check не логируем
		if r.URL.Path == "/health_check" {
			next.ServeHTTP(w, r)
			return
		}

		// Заголовки
		reqHeadersStr := fmt.Sprintf("%v", r.Header)
		s.log.Info(
			"Request received",
			zap.String("method", r.Method),
			zap.String("route", r.URL.Path),
			zap.String("ip", r.RemoteAddr),
			zap.String("headers", reqHeadersStr[4:len(reqHeadersStr)-1]),
		)

		// Тело
		body, _ := io.ReadAll(r.Body)
		s.log.Info(string(body))
		// важно: восстанавливаем тело, иначе хендлер не увидит его
		r.Body = io.NopCloser(strings.NewReader(string(body)))

		next.ServeHTTP(w, r)
	})
}
