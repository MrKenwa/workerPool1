package workerPoolUC

import "workerPool1/internal/entity"

func (wp *WorkerPool) GetStatuses() map[string]entity.Status {
	res := make(map[string]entity.Status)
	wp.mu.Lock()
	defer wp.mu.Unlock()

	for key, value := range wp.statuses {
		res[key] = value
	}

	return res
}
