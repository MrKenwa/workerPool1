package worker_pool

import (
	"encoding/json"
	"errors"
	"net/http"

	"workerPool1/internal/entity"
)

var (
	ErrMissingFields = errors.New("missing fields or invalid max_retries")
)

type EnqueueRequest struct {
	ID         string `json:"id"`
	Payload    string `json:"payload"`
	MaxRetries int    `json:"max_retries"`
}

func (h *Handler) Enqueue(w http.ResponseWriter, r *http.Request) {
	var req EnqueueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, err, http.StatusBadRequest)
		return
	}

	if req.ID == "" || req.Payload == "" || req.MaxRetries < 0 {
		writeJSONError(w, ErrMissingFields, http.StatusBadRequest)
		return
	}

	err := h.service.Enqueue(mapEnqueueRequest(req))
	if err != nil {
		writeJSONError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(map[string]string{
		"data": "successful enqueued task",
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func mapEnqueueRequest(req EnqueueRequest) entity.Task {
	return entity.Task{
		ID:         req.ID,
		Payload:    req.Payload,
		MaxRetries: req.MaxRetries,
		Attempts:   0,
	}
}

func writeJSONError(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	resp := map[string]string{
		"error": err.Error(),
	}

	if encodeErr := json.NewEncoder(w).Encode(resp); encodeErr != nil {
		http.Error(w, err.Error(), code)
	}
}
