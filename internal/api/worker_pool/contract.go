package worker_pool

import (
	"workerPool1/internal/entity"
)

type (
	service interface {
		Enqueue(task entity.Task) error
		GetStatuses() map[string]entity.Status
	}
)
