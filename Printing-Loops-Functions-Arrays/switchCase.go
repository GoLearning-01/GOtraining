package main

import "fmt"

func main() {

	fmt.Println("Enter a number as string: ")
	var stringNum string
	fmt.Scanln(&stringNum)

	switch stringNum {
	case "one":
		fmt.Println("You typed 1")
	case "two":
		fmt.Println("You typed 2")
	case "three":
		fmt.Println("You typed 3")
	default:
		fmt.Println("Out of range!")
	}

}
