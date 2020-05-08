package main

import (
	"fmt"
)

func rectAP(x, y float64) (area float64, perimeter float64) {
	area = x * y
	perimeter = 2 * (x + y)
	return
}

func main() {

	fmt.Println("Enter length of a rectangle")
	var length float64
	fmt.Scanln(&length)

	fmt.Println("Enter bredth of a rectangle")
	var bredth float64
	fmt.Scanln(&bredth)

	a, b := rectAP(length, bredth)
	fmt.Println("Area =", a)
	fmt.Println("Perimeter=", b)

}
