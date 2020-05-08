package main

import (
	"fmt"
	"math"
)

func maximumValue(x, y float64) (maximum float64) {
	maximum = math.Max(x, y)
	return
}

func main() {

	array := [10]float64{2, 3, 5, 23, 12, 67, 34, 23, 12, 5}

	maxVal := 0.0
	for i := 0; i < len(array); i++ {
		maxVal = maximumValue(maxVal, array[i])
	}

	array2 := array

	for i := 0; i < len(array2); i++ {
		if array2[i] == maxVal {
			array2[i] = 0
		}
	}

	maxVal = 0.0
	for i := 0; i < len(array2); i++ {
		maxVal = maximumValue(maxVal, array2[i])
	}

	fmt.Println("Second Largest Number =", maxVal)

}
