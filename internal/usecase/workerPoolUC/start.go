package workerPoolUC

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const maxDelay = 30 * time.Second

func (wp *WorkerPool) Start() {
	wg := sync.WaitGroup{}
	wg.Add(wp.cfg.Queue.Workers)

	for i := 0; i < wp.cfg.Queue.Workers; i++ {
		go func() {
			defer wg.Done()
			wp.startWorker()
		}()
	}

	wg.Wait()
	close(wp.doneCh)
}

func (wp *WorkerPool) startWorker() {
	for task := range wp.taskCh {
		log.Printf("Start processing task %s", task.ID)
		attempt := 0
		for ; attempt <= task.MaxRetries; attempt++ { // <= чтобы была минимум одна попытка
			err := wp.processTask(task)
			if err != nil {
				// экспоненциальный бэкофф с джиттером
				backoff := time.Duration((1<<attempt)*wp.cfg.Queue.BaseBackoff) * time.Millisecond
				if backoff > maxDelay {
					backoff = maxDelay
				}

				jitter := time.Duration(rand.Int63n(int64(backoff)))
				log.Printf("Error in process task %s, sleep %v", task.ID, backoff+jitter)
				time.Sleep(backoff + jitter)
			} else {
				break
			}
		}
		log.Printf("Task %s was processed successful after %d attempts", task.ID, attempt)
	}
}
