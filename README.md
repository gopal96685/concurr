# Worker Pool

This package provides an implementation of a worker pool pattern for executing tasks concurrently using a pool of worker goroutines. It allows you to submit tasks to the worker pool and receive results as tasks are completed.

## Usage

1. Import the package:

   ```go
   import "github.com/your-username/workerpool"

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

    type MyTask struct {
        Value int
    }

    func (t MyTask) Execute() (workerpool.Result, error) {
        // Perform task logic
        // ...
        return MyResult{Value: t.Value * 2}, nil
    }

    for i := 0; i < 100; i++ {
        task := MyTask{Value: i}
        if err := pool.Submit(task); err != nil {
            // Handle error
        }
    }

    for result := range pool.Results() {
        // Process result
        fmt.Println(result.Value())
    }

    pool.Shutdown()
```

In this implementation, the functions have low complexity and focus on performing their specific tasks efficiently. The `worker` struct and its associated functions (`newWorker`, `start`, and `stop`) handle the execution of tasks in a concurrent and non-blocking manner.

The `WorkerPool` struct encapsulates the worker management, task queueing, and result handling. The `NewWorkerPool` function creates a new worker pool instance with the given configuration, while the `Submit` method allows users to submit tasks for execution.

The `Results` method returns a read-only channel that receives results as tasks are completed, and the `Shutdown` method gracefully stops the worker pool and waits for all tasks to complete.

The `Config` struct represents the configuration options for the worker pool, such as the number of workers, task queue size, and idle timeout.

The `Task` and `Result` interfaces define the contracts for tasks and results, respectively, allowing users to implement their own task logic and result types.

Additionally, the `errors.go` file contains predefined error values for common error scenarios, and the `doc.go` file provides package-level documentation with usage examples.