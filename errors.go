package concurr

import "errors"

var (
	ErrPoolClosed    = errors.New("worker pool is closed")
	ErrTaskQueueFull = errors.New("task queue is full")
)
