package workerPoolHandler

import (
	"encoding/json"
	"net/http"
	"workerPool1/internal/entity"
)

type EnqueueRequest struct {
	ID         string `json:"id"`
	Payload    string `json:"payload"`
	MaxRetries int    `json:"max_retries"`
}

func (h *WorkerPoolHandler) Enqueue(w http.ResponseWriter, r *http.Request) {
	var req EnqueueRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.ID == "" || req.Payload == "" || req.MaxRetries < 0 {
		http.Error(w, "missing fields or invalid max_retries", http.StatusBadRequest)
		return
	}

	err := h.workerPoolUC.Enqueue(mapEnqueueRequest(req))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"data": "successful enqueued task",
	})
}

func mapEnqueueRequest(req EnqueueRequest) *entity.Task {
	return &entity.Task{
		ID:         req.ID,
		Payload:    req.Payload,
		MaxRetries: req.MaxRetries,
		Attempts:   0,
	}
}
