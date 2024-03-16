/*
Package workerpool provides an implementation of a worker pool pattern for
executing tasks concurrently using a pool of worker goroutines.

Usage:

	// Create a new worker pool with a configuration
	config := workerpool.Config{
	    NumWorkers:      10,
	    TaskQueueSize:   100,
	    TaskQueueBuffer: 10,
	    IdleTimeout:     time.Minute,
	}
	pool, err := workerpool.NewWorkerPool(config)
	if err != nil {
	    // Handle error
	}

	// Submit tasks to the worker pool
	for i := 0; i < 100; i++ {
	    task := MyTask{Value: i}
	    if err := pool.Submit(task); err != nil {
	        // Handle error
	    }
	}

	// Receive results from the worker pool
	for result := range pool.Results() {
	    // Process result
	}

	// Shutdown the worker pool gracefully
	pool.Shutdown()
*/
package concurr
