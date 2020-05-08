package main

import (
	"fmt"
)

func add(x float64, y float64) float64 {
	return x + y
}

func sub(x float64, y float64) float64 {
	return x - y
}

func mul(x float64, y float64) float64 {
	return x * y
}

func div(x float64, y float64) float64 {
	return x / y
}

func per(x float64, y float64) float64 {
	return (x / y) * 100
}

func main() {

	var selection int
	var choice string

top:

	fmt.Println("-----------------------")
	fmt.Println("    ---  Menu  ---")
	fmt.Println("Please make a selection")
	fmt.Println("-----------------------")
	fmt.Println("   1. Addition")
	fmt.Println("   2. Subtraction")
	fmt.Println("   3. Multiplication")
	fmt.Println("   4. Division")
	fmt.Println("   5. Class Percentage")
	fmt.Println("   6. Number Series")
	fmt.Println("-----------------------")
	fmt.Println("-----------------------")

	fmt.Scanln(&selection)

	switch selection {
	case 1:
		fmt.Println("Enter first number:")
		var x float64
		fmt.Scanln(&x)

		fmt.Println("Enter second number:")
		var y float64
		fmt.Scanln(&y)

		fmt.Println("-----------------------")
		fmt.Println("Sum = ", add(x, y))

	case 2:
		fmt.Println("Enter first number:")
		var x float64
		fmt.Scanln(&x)

		fmt.Println("Enter second number:")
		var y float64
		fmt.Scanln(&y)

		fmt.Println("-----------------------")
		fmt.Println("Difference = ", sub(x, y))

	case 3:
		fmt.Println("Enter first number:")
		var x float64
		fmt.Scanln(&x)

		fmt.Println("Enter second number:")
		var y float64
		fmt.Scanln(&y)

		fmt.Println("-----------------------")
		fmt.Println("Product = ", mul(x, y))

	case 4:
		fmt.Println("Enter first number:")
		var x float64
		fmt.Scanln(&x)

		fmt.Println("Enter second number:")
		var y float64
		fmt.Scanln(&y)

		fmt.Println("-----------------------")
		fmt.Println("Quotient = ", div(x, y))

	case 5:
		fmt.Println("Enter your score:")
		var x float64
		fmt.Scanln(&x)

		fmt.Println("Enter total score:")
		var y float64
		fmt.Scanln(&y)

		fmt.Println("-----------------------")
		fmt.Println("Percentage = ", per(x, y), "%")

	case 6:
		fmt.Println("Enter starting number:")
		var x float64
		fmt.Scanln(&x)

		fmt.Println("Enter ending number:")
		var y float64
		fmt.Scanln(&y)

		fmt.Println("-----------------------")
		for x <= y {
			fmt.Println(x)
			x++
		}

	default:
		fmt.Println("Error")

	}

	fmt.Println("-----------------------")
	fmt.Println("More calculations(y/n)?")
	fmt.Scanln(&choice)
	fmt.Println("-----------------------")

	if choice == "y" {
		goto top
	}

}
