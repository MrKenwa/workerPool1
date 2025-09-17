package worker_pool

import (
	"workerPool1/internal/entity"
)

func (wp *Service) changeTaskStatus(id string, status entity.Status) {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	wp.statuses[id] = status
}
