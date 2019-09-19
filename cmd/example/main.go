package main

import (
	"fmt"

	"github.com/mr-panta/demo-calculator"
)

func main() {
	cal := calculator.NewCalculator(100)
	cal.AddWith(10)
	cal.SubWith(5)
	cal.DivBy(12)
	cal.MulBy(7)
	fmt.Println(cal.GetResult())
}
