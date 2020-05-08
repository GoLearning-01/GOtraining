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

	fmt.Println("")
	fmt.Println("**************************")
	fmt.Println("Array Loaded Successfully!")
	fmt.Println("**************************")
	fmt.Println("")

	fmt.Println("Enter a search value:")
	var searchVal int
	fmt.Scanln(&searchVal)

	fmt.Println("")

	testValue := false
	indexValue := 0

	for i := 0; i < len(array); i++ {
		if array[i] == searchVal {
			testValue = true
			indexValue = i
		}
	}

	if testValue == true {
		fmt.Println("Element Present! Index =", indexValue)
	} else {
		fmt.Println("Element is not present in array!")
	}

	fmt.Println("")
}
