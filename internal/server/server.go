package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	workerPoolHandler "workerPool1/internal/api/worker_pool"
	"workerPool1/internal/config"
	workerPoolService "workerPool1/internal/service/worker_pool"
)

type Server struct {
	srv *http.Server
	mux *http.ServeMux
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	mux := &http.ServeMux{}
	return &Server{
		srv: &http.Server{
			Addr:    cfg.ServerConfig.Host,
			Handler: mux,
		},
		mux: mux,
		cfg: cfg,
	}
}

func (s *Server) Start() error {
	workerPool, err := workerPoolService.New(&s.cfg.QueueConfig)
	if err != nil {
		log.Printf("error while creating worker pool: %v", err)
		return err
	}
	go workerPool.Start()

	wpHandler := workerPoolHandler.New(workerPool)
	workerPoolHandler.MapRoutes(s.mux, wpHandler)

	s.mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	go func() {
		log.Printf("Server started on %s", s.cfg.ServerConfig.Host)
		if err := s.srv.ListenAndServe(); err != nil {
			log.Printf("Listen error: %v", err)
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down...")

	if err = workerPool.Close(); err != nil {
		log.Printf("error while closing worker pool: %v", err)
	}
	log.Print("Worker pool is closed")

	if err = s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("server shutdown error: %v", err)
		return err
	}
	log.Print("Server is stopped")

	return nil
}
