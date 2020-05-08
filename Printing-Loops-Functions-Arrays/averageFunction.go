// Average using Function

package main

import "fmt"

func averageFunction(x, y float64) (average float64) {
	average = (x + y) / 2
	return
}

func main() {

	fmt.Println("Enter first number:")
	var firstNum float64
	fmt.Scanln(&firstNum)

	fmt.Println("Enter second number:")
	var secondNum float64
	fmt.Scanln(&secondNum)

	fmt.Println("The average is", averageFunction(firstNum, secondNum))
}
