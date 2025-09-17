package worker_pool

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) Statuses(w http.ResponseWriter, _ *http.Request) {
	res := h.service.GetStatuses()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"data": res,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
