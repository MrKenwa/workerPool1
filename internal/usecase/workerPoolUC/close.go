package workerPoolUC

import "errors"

func (wp *WorkerPool) Close() error {
	// Локаем мьютекс чтобы атомарно закрыть канал и установить флаг
	wp.mu.Lock()
	if wp.isClosed {
		return errors.New("worker pool is already closed")
	}

	// Отмечаем очередь закрытой и закрываем канал тасок
	wp.isClosed = true
	close(wp.taskCh)
	wp.mu.Unlock()

	// Ждем завершения всех воркеров
	<-wp.doneCh

	return nil
}
