package main

import "fmt"

func main() {

	array := [11]int{-1, 2, -3, -4, -5, 6, 0, 1, -2, 3, 4}

	negative := 0

	for i := 0; i < len(array); i++ {
		if array[i] < 0 {
			negative = negative + 1
		}
	}
	fmt.Println("Total Negative =", negative)
}
