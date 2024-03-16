package concurr

// Task represents a unit of work that can be executed by a worker.
type Task interface {
	// Execute performs the task logic and returns a result.
	Execute() (Result, error)
}
