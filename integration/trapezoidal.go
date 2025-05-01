package integration

import (
	"math"

	"github.com/karmek-k/go-integration/functions"
)

// TrapezoidalIntegral is an Integral calculated using the trapezoidal rule.
type TrapezoidalIntegral struct {
	Function functions.Function
	Cuts float64
}

func (i *TrapezoidalIntegral) Calculate(a, b float64) float64 {
	// distance between each trapezoid
	dx := math.Abs(a - b) / i.Cuts

	result := 0.0

	// assuming that a < b
	for x := a; x < b - dx; x += dx {
		y1 := i.Function.Evaluate(x)
		y2 := i.Function.Evaluate(x + dx)

		// trapezoid surface
		result += 0.5 * (y1 + y2) * dx 
	}

	return result
}
