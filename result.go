package concurr

// Result represents the output of a task execution.
type Result interface {
	// Value returns the result value.
	Value() interface{}
}
