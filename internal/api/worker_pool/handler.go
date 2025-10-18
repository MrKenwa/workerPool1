package worker_pool

import (
	"net/http"

	"workerPool1/internal/logger"
)

type Handler struct {
	service service
	log     *logger.Logger
}

func New(service service, log *logger.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func MapRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /enqueue", h.Enqueue)
	mux.HandleFunc("GET /statuses", h.Statuses)
}
