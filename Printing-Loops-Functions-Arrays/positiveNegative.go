package main

import "fmt"

func main() {

	fmt.Println("Enter a number:")
	var num int
	fmt.Scanln(&num)

	if num >= 0 {
		fmt.Println("The number is Positive")
	} else {
		fmt.Println("The number is Negative")
	}

}
