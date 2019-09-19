package calculator

type operator int

const (
	operatorAdd operator = iota + 1
	operatorSub
	operatorDiv
	operatorMul
)

// Calculator contains basic math operation functions.
type Calculator interface {
	AddWith(b float64)
	SubWith(b float64)
	DivBy(b float64)
	MulBy(b float64)
	GetResult() (result float64)
}

type defaultCalculator struct {
	a float64
}
