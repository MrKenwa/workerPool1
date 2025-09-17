package workerPoolUC

import (
	"errors"
	"math/rand"
	"time"

	"workerPool1/internal/entity"
)

func (wp *WorkerPool) processTask(task *entity.Task) error {
	wp.changeTaskStatus(task.ID, entity.StatusRunning)
	// Симулируем обработку
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
	time.Sleep(delay)

	if rand.Intn(100) < 20 { // 20% ошибка
		wp.changeTaskStatus(task.ID, entity.StatusFailed)
		return errors.New("error while startWorker task")
	} else {
		wp.changeTaskStatus(task.ID, entity.StatusDone)
		return nil
	}
}
