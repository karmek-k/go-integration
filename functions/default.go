package functions

import "math"

// GetDefault provides functions available by default.
func GetDefault() []Function {
	return []Function{
		{
			Name:     "sin(x)",
			Domain:   "R",
			Evaluate: math.Sin,
		},
		{
			Name:     "0.5x^2 + 7x - 3.5",
			Domain:   "R",
			Evaluate: quadratic,
		},
		{
			Name:     "sqrt(x)",
			Domain:   "non-negative R",
			Evaluate: math.Sqrt,
		},
	}
}

// quadratic represents a certain quadraticc Function's Evaluator.
func quadratic(x float64) float64 {
	const A float64 = 0.5
	const B float64 = 7.0
	const C float64 = -3.5

	return A*(math.Pow(x, 2.0)) + B*x + C
}
