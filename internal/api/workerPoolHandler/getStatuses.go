package workerPoolHandler

import (
	"encoding/json"
	"net/http"
)

func (h *WorkerPoolHandler) Statuses(w http.ResponseWriter, r *http.Request) {
	res := h.workerPoolUC.GetStatuses()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
