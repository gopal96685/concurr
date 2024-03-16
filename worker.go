package concurr

// worker represents a single worker goroutine within the pool.
type worker struct {
	id     int
	pool   *WorkerPool
	taskCh chan Task
	stopCh chan struct{}
}

func newWorker(id int, pool *WorkerPool) *worker {
	return &worker{
		id:     id,
		pool:   pool,
		taskCh: pool.taskQueue,
		stopCh: pool.stopCh,
	}
}

func (w *worker) start() {
	go func() {
		defer w.pool.wg.Done()
		for {
			select {
			case task, ok := <-w.taskCh:
				if !ok {
					return
				}
				result, err := task.Execute()
				if err != nil {
					// Handle error
				}
				w.pool.resultChan <- result
			case <-w.stopCh:
				return
			}
		}
	}()
}

func (w *worker) stop() {
	close(w.stopCh)
}
