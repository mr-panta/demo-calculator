package calculator_test

import (
	"testing"

	"github.com/mr-panta/demo-calculator"
)

func TestAdd(t *testing.T) {
	cal := calculator.NewCalculator(100)
	cal.AddWith(50)
	if cal.GetResult() != 150 {
		t.Errorf("adder function is wrong, expected: %f, actual: %f", 150.0, cal.GetResult())
	}
}

func TestSub(t *testing.T) {
	cal := calculator.NewCalculator(100)
	cal.SubWith(50)
	if cal.GetResult() != 50 {
		t.Errorf("substractor function is wrong, expected: %f, actual: %f", 50.0, cal.GetResult())
	}
}

func TestDiv(t *testing.T) {
	cal := calculator.NewCalculator(100)
	cal.DivBy(50)
	if cal.GetResult() != 2 {
		t.Errorf("divider function is wrong, expected: %f, actual: %f", 2.0, cal.GetResult())
	}
}

func TestMul(t *testing.T) {
	cal := calculator.NewCalculator(100)
	cal.MulBy(50)
	if cal.GetResult() != 5000 {
		t.Errorf("multiplier function is wrong, expected: %f, actual: %f", 5000.0, cal.GetResult())
	}
}
