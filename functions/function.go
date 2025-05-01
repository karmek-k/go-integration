package functions

type Evaluator func(float64)float64

type Function struct {
	Name string
	Domain string
	Evaluate Evaluator
}
