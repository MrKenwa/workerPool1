package workerPoolUC

import (
	"errors"

	"workerPool1/internal/entity"
)

func (wp *WorkerPool) Enqueue(task *entity.Task) error {
	if task == nil {
		return errors.New("task is nil")
	}

	wp.mu.Lock()
	defer wp.mu.Unlock()
	if wp.isClosed {
		return errors.New("queue is closed")
	}

	select {
	case wp.taskCh <- task:
		wp.statuses[task.ID] = entity.StatusQueued
		return nil
	default:
		return errors.New("queue is overcrowded") // очередь переполнена
	}
}
