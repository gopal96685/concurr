package concurr

import "sync"

// WorkerPool represents a pool of worker goroutines that can execute tasks concurrently.
type WorkerPool struct {
	config     Config
	taskQueue  chan Task
	resultChan chan Result
	stopCh     chan struct{}
	wg         sync.WaitGroup
	workers    []*worker
}

// NewWorkerPool creates a new worker pool with the given configuration.
func NewWorkerPool(config Config) (*WorkerPool, error) {
	// Validate and apply configuration
	NewPool := WorkerPool{
		config: Config{
			NumWorkers:      config.NumWorkers,
			TaskQueueSize:   config.TaskQueueSize,
			TaskQueueBuffer: config.TaskQueueBuffer,
			IdleTimeout:     config.IdleTimeout,
		},
		taskQueue:  make(chan Task, config.TaskQueueSize),
		resultChan: make(chan Result),
		stopCh:     make(chan struct{}),
		wg:         sync.WaitGroup{},
		workers:    make([]*worker, config.NumWorkers),
	}

	// Initialize worker pool struct
	for i := 0; i < config.NumWorkers; i++ {
		NewWorker := newWorker(i, &NewPool)
		NewPool.workers = append(NewPool.workers, NewWorker)
	}

	// Start workers
	for i := 0; i < config.NumWorkers; i++ {
		NewPool.wg.Add(1)
		NewPool.workers[i].start()
	}

	return &NewPool, nil
}

// Submit submits a task to the worker pool for execution.
func (wp *WorkerPool) Submit(task Task) error {
	//validate the task

	// Submit task to task queue
	wp.taskQueue <- task
	return nil
}

// Results returns a channel that receives results as tasks are completed.
func (wp *WorkerPool) Results() <-chan Result {
	// Return a read-only result channel
	return wp.resultChan
}

// Shutdown gracefully shuts down the worker pool and waits for all tasks to complete.
func (wp *WorkerPool) Shutdown() {
	// Signal workers to stop
	for _, worker := range wp.workers {
		worker.stop()
		wp.wg.Done()
	}

	// Wait for all workers to finish
	wp.wg.Wait()

	// Close result channel
	close(wp.resultChan)
}
