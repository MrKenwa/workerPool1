package workerPoolHandler

import (
	"net/http"

	"workerPool1/internal/entity"
)

type (
	WorkerPoolUC interface {
		Enqueue(task *entity.Task) error
		GetStatuses() map[string]entity.Status
	}

	Handlers interface {
		Enqueue(w http.ResponseWriter, r *http.Request)
		Statuses(w http.ResponseWriter, r *http.Request)
	}
)

type WorkerPoolHandler struct {
	workerPoolUC WorkerPoolUC
}

func New(workerPoolUC WorkerPoolUC) *WorkerPoolHandler {
	return &WorkerPoolHandler{
		workerPoolUC: workerPoolUC,
	}
}

func MapRoutes(mux *http.ServeMux, h Handlers) {
	mux.HandleFunc("POST /enqueue", h.Enqueue)
	mux.HandleFunc("GET /statuses", h.Statuses)
}
