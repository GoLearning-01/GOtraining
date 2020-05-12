package main

import "fmt"

func swapValue(m, n *int) (*int, *int) {
	tempA := *m
	tempB := *n
	*m = tempB
	*n = tempA
	return m, n
}

func main() {

	a, b := 6, 12

	p1 := &a
	p2 := &b

	fmt.Println(p1, p2)

	x, y := swapValue(p1, p2)

	a = *x
	b = *y

	fmt.Println(a, b)

}
