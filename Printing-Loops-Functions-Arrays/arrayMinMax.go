package main

import (
	"fmt"
	"math"
)

func minValue(x, y float64) (min float64) {
	min = math.Min(x, y)
	return
}

func maxValue(x, y float64) (max float64) {
	max = math.Max(x, y)
	return
}

func main() {

	a1 := [6]float64{8, 3, 2, 45, 6, 7}

	minimum := a1[0]
	for i := 1; i < len(a1); i++ {
		minimum = minValue(minimum, a1[i])
	}

	maximum := a1[0]
	for i := 1; i < len(a1); i++ {
		maximum = maxValue(maximum, a1[i])

	}

	fmt.Println("Minimum Value = ", minimum)
	fmt.Println("Minimum Value = ", maximum)

}
