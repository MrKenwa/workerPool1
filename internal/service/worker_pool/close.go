package worker_pool

import "errors"

var (
	ErrWorkerPoolAlreadyClosed = errors.New("worker pool is already closed")
)

func (wp *Service) Close() error {
	// Локаем мьютекс чтобы атомарно закрыть канал и установить флаг
	wp.mu.Lock()
	if wp.isClosed {
		return ErrWorkerPoolAlreadyClosed
	}

	// Отмечаем очередь закрытой и закрываем канал тасок
	wp.isClosed = true
	close(wp.taskCh)
	wp.mu.Unlock()

	// Ждем завершения всех воркеров
	<-wp.doneCh

	return nil
}
