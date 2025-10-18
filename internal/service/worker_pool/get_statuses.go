package worker_pool

import "workerPool1/internal/entity"

func (wp *Service) GetStatuses() map[string]entity.Status {
	res := make(map[string]entity.Status)
	wp.mu.Lock()
	defer wp.mu.Unlock()

	for key, value := range wp.statuses {
		res[key] = value
	}

	return res
}
