package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Enter first number:")
	var firstNum float64
	fmt.Scanln(&firstNum)

	fmt.Println("Enter second number")
	var secondNum float64
	fmt.Scanln(&secondNum)

	var greaterNum = math.Max(firstNum, secondNum)

	fmt.Println("The bigger number is:", greaterNum)

}
