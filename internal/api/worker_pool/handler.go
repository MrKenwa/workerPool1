package worker_pool

import (
	"net/http"
)

type Handler struct {
	service service
}

func New(service service) *Handler {
	return &Handler{
		service: service,
	}
}

func MapRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /enqueue", h.Enqueue)
	mux.HandleFunc("GET /statuses", h.Statuses)
}
