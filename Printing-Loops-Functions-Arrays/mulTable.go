package main

import "fmt"

func main() {

	fmt.Println("Enter a number for Multiplication Table:")
	var number int
	fmt.Scanln(&number)

	for i := 1; i <= 10; i++ {
		fmt.Println(number, "X", i, "=", number*i)
	}
}
