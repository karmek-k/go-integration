package functions

// Evaluator calculates a single value from a single argument.
type Evaluator func(float64)float64

// Function represents a single-variable real function.
// Symbolically f : X -> R, where X is a subset of R.
type Function struct {
	Name string
	Domain string
	Evaluate Evaluator
}
