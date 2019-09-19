package calculator

func NewCalculator(a float64) Calculator {
	return &defaultCalculator{a: a}
}

func (c *defaultCalculator) calculate(op operator, a, b float64) (result float64) {
	switch op {
	case operatorAdd:
		result = a + b
	case operatorSub:
		result = a - b
	case operatorDiv:
		result = a / b
	case operatorMul:
		result = a * b
	}
	return result
}

// AddWith is used to increase result in Calculator by b.
func (c *defaultCalculator) AddWith(b float64) {
	c.a = c.calculate(operatorAdd, c.a, b)
}

// SubWith is used to decrease result in Calculator by b.
func (c *defaultCalculator) SubWith(b float64) {
	c.a = c.calculate(operatorSub, c.a, b)
}

// DivBy is used to divide result in Calculator by b.
func (c *defaultCalculator) DivBy(b float64) {
	c.a = c.calculate(operatorDiv, c.a, b)
}

// MulBy is used to multiply result in Calculator by b.
func (c *defaultCalculator) MulBy(b float64) {
	c.a = c.calculate(operatorMul, c.a, b)
}

// GetResult returns result in Calculator.
func (c *defaultCalculator) GetResult() (result float64) {
	return c.a
}
