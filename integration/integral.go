package integration

// Integral represents a single definite integral.
type Integral interface {
	// Calculate finds an approximate value of the integral.
	Calculate(a float64, b float64) float64
}
