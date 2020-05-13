package main

import (
	"errors"
	"fmt"
)

func Checker(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("Enter positive number")
	}
	return n, nil
}

func main() {
	fmt.Println("Enter a number:")
	var num int
	fmt.Scanln(&num)

	posNum, err := Checker(num)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Number", posNum, "is positive!")
}
