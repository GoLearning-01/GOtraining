package main

import "fmt"

func main() {

	// Taking user input for array

	fmt.Println("Enter first number:")
	var a int
	fmt.Scanln(&a)

	fmt.Println("Enter second number:")
	var b int
	fmt.Scanln(&b)

	fmt.Println("Enter third number:")
	var c int
	fmt.Scanln(&c)

	fmt.Println("Enter fourth number:")
	var d int
	fmt.Scanln(&d)

	fmt.Println("Enter fifth number:")
	var e int
	fmt.Scanln(&e)

	// User input finish

	array := [...]int{a, b, c, d, e}

	var sum int = 0

	for i := 0; i < len(array); i++ {
		sum = sum + array[i]
	}

	fmt.Println("Sum = ", sum)

}
