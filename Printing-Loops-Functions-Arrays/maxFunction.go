// Maximum of two numbers using finction

package main

import (
	"fmt"
	"math"
)

func maxNum(x, y float64) float64 {
	return math.Max(x, y)
}

func main() {

	fmt.Println("Enter first number:")
	var firstNum float64
	fmt.Scanln(&firstNum)

	fmt.Println("Enter second number:")
	var secondNum float64
	fmt.Scanln(&secondNum)

	fmt.Println("The maximum number is", maxNum(firstNum, secondNum))

}
