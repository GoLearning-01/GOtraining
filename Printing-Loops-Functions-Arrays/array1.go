package main

import "fmt"

func main() {
	var x [2][3]int
	fmt.Println(x)

	var y [5]complex128
	fmt.Println(y)

	z := [...]int{100, 150, 200, 250, 300}

	for i := 0; i < len(z); i++ {
		fmt.Println("x[", i, "] = ", z[i])
	}
}
