package worker_pool

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"

	"workerPool1/internal/entity"
)

var ErrTaskProcessing = errors.New("error while process task")

const maxDelay = 30 * time.Second

func (wp *Service) Start() {
	wg := sync.WaitGroup{}
	wg.Add(wp.cfg.Workers)

	for i := 0; i < wp.cfg.Workers; i++ {
		go func() {
			defer wg.Done()
			wp.startWorker()
		}()
	}

	wg.Wait()
	close(wp.doneCh)
}

func (wp *Service) startWorker() {
	for task := range wp.taskCh {
		log.Printf("Start processing task %s", task.ID)
		attempt := 0
		for ; attempt <= task.MaxRetries; attempt++ { // <= чтобы была минимум одна попытка
			err := wp.processTask(task)
			if err == nil {
				break
			}

			// экспоненциальный бэкофф с джиттером
			backoff := wp.cfg.BaseBackoff * time.Duration(math.Pow(2, float64(attempt)))
			if backoff > maxDelay {
				backoff = maxDelay
			}

			jitter := time.Duration(rand.Int63n(int64(backoff)))
			log.Printf("Error in process task %s, sleep %v", task.ID, backoff+jitter)
			time.Sleep(backoff + jitter)
		}
		log.Printf("Task %s was processed successful after %d attempts", task.ID, attempt)
	}
}

func (wp *Service) processTask(task entity.Task) error {
	wp.changeTaskStatus(task.ID, entity.StatusRunning)
	// Симулируем обработку
	delay := time.Duration(100+rand.Intn(400)) * time.Millisecond
	time.Sleep(delay)

	if rand.Intn(100) < 20 { // 20% ошибка
		wp.changeTaskStatus(task.ID, entity.StatusFailed)
		return ErrTaskProcessing
	} else {
		wp.changeTaskStatus(task.ID, entity.StatusDone)
		return nil
	}
}
