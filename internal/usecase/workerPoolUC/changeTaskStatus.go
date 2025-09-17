package workerPoolUC

import (
	"workerPool1/internal/entity"
)

func (wp *WorkerPool) changeTaskStatus(id string, status entity.Status) {
	wp.mu.Lock()
	wp.statuses[id] = status
	wp.mu.Unlock()
}
