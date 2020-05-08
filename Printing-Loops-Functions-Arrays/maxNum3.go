package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Enter first number:")
	var first float64
	fmt.Scanln(&first)

	fmt.Println("Enter second number:")
	var second float64
	fmt.Scanln(&second)

	fmt.Println("Enter third number:")
	var third float64
	fmt.Scanln(&third)

	var maxNum float64 = math.Max(first, second)
	var maxNum3 float64 = math.Max(maxNum, third)

	fmt.Println("Maximum number is:", maxNum3)
}
