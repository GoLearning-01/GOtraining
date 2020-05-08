// Accurate under 100000

package main

import (
	"fmt"
	"math"
)

func minimumValue(x, y float64) (minimum float64) {
	minimum = math.Min(x, y)
	return
}

func main() {

	array := [10]float64{2, 3, 5, 23, 12, 67, 34, 23, 12, 5}

	minVal := 99999.9
	for i := 0; i < len(array); i++ {
		minVal = minimumValue(minVal, array[i])
	}

	array2 := array

	for i := 0; i < len(array2); i++ {
		if array2[i] == minVal {
			array2[i] = 999
		}
	}

	minVal = 99999.9
	for i := 0; i < len(array2); i++ {
		minVal = minimumValue(minVal, array2[i])
	}

	fmt.Println("Second Smallest Number =", minVal)

}
