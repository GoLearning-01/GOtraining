package main

import "fmt"

func main() {

	a, b := 12, 16

	p := &a         // Setting p as a address of a
	fmt.Println(p)  // returns memory address
	fmt.Println(*p) // returns value

	p = &b // setting p as a address of b now
	fmt.Println(p)
	fmt.Println(*p)

	fmt.Println(a, b)
}
