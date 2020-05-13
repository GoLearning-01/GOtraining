/*
Calculate area of a rectangle.
Pass error if length or bredth is negative
*/

package main

import (
	"errors"
	"fmt"
)

func area(l, b float64) (float64, error) {
	if l < 0 || b < 0 {
		return 0, errors.New("Error: Value can not be less than 0")
	} else {
		return l * b, nil
	}
}

func main() {

	fmt.Println("Enter length:")
	var length float64
	fmt.Scanln(&length)

	fmt.Println("Enter bredth:")
	var bredth float64
	fmt.Scanln(&bredth)

	result, err := area(length, bredth)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Area =", result)
	}

}
