package workerPoolUC

import (
	"errors"
	"sync"

	"workerPool1/internal/config"
	"workerPool1/internal/entity"
)

type WorkerPool struct {
	cfg      *config.Config
	mu       sync.Mutex
	statuses map[string]entity.Status
	isClosed bool
	taskCh   chan *entity.Task
	doneCh   chan struct{}
}

func NewWorkerPool(cfg *config.Config) (*WorkerPool, error) {
	if cfg.Queue.Workers <= 0 {
		return nil, errors.New("workersCount num is incorrect")
	}

	if cfg.Queue.QueueSize <= 0 {
		return nil, errors.New("queueSize num is incorrect")
	}

	return &WorkerPool{
		cfg:      cfg,
		mu:       sync.Mutex{},
		statuses: make(map[string]entity.Status),
		isClosed: false,
		taskCh:   make(chan *entity.Task, cfg.Queue.QueueSize),
		doneCh:   make(chan struct{}),
	}, nil
}
