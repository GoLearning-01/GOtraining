package main

import "fmt"

func main() {

	fmt.Println("Press 1 for area of Square and 2 for area of Rectangle")
	var choice int
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println("Enter length of a side:")
		var lengthSquare int
		fmt.Scanln(&lengthSquare)

		fmt.Println("Area of square is:", lengthSquare*lengthSquare)
	} else {
		fmt.Println("Enter length of rectangle:")
		var lengthRectangle int
		fmt.Scanln(&lengthRectangle)

		fmt.Println("Enter bredth of rectangle:")
		var bredthRectangle int
		fmt.Scanln(&bredthRectangle)

		fmt.Println("Area of rectangle is:", lengthRectangle*bredthRectangle)
	}

}
