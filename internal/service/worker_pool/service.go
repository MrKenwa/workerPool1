package worker_pool

import (
	"errors"
	"sync"

	"workerPool1/internal/config"
	"workerPool1/internal/entity"
)

var (
	ErrWorkersCountBelowZero = errors.New("the number of workers is less than zero")
	ErrQueueSizeBelowZero    = errors.New("the queue size is less than zero")
)

type Service struct {
	cfg      *config.QueueConfig
	mu       sync.Mutex
	statuses map[string]entity.Status
	taskCh   chan entity.Task
	doneCh   chan struct{}
	isClosed bool
}

func New(cfg *config.QueueConfig) (*Service, error) {
	if cfg.Workers <= 0 {
		return nil, ErrWorkersCountBelowZero
	}

	if cfg.QueueSize <= 0 {
		return nil, ErrQueueSizeBelowZero
	}

	return &Service{
		cfg:      cfg,
		mu:       sync.Mutex{},
		statuses: make(map[string]entity.Status),
		isClosed: false,
		taskCh:   make(chan entity.Task, cfg.QueueSize),
		doneCh:   make(chan struct{}),
	}, nil
}
