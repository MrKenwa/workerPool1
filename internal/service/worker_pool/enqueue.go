package worker_pool

import (
	"errors"

	"workerPool1/internal/entity"
)

var (
	ErrQueueIsClosed      = errors.New("queue is closed")
	ErrQueueIsOvercrowded = errors.New("queue is overcrowded")
)

func (wp *Service) Enqueue(task entity.Task) error {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	if wp.isClosed {
		return ErrQueueIsClosed
	}

	select {
	case wp.taskCh <- task:
		wp.statuses[task.ID] = entity.StatusQueued
		return nil
	default:
		return ErrQueueIsOvercrowded // очередь переполнена
	}
}
