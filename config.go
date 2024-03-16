package concurr

import "time"

// Config represents the configuration options for the worker pool.
type Config struct {
	NumWorkers      int
	TaskQueueSize   int
	TaskQueueBuffer int
	IdleTimeout     time.Duration
}
