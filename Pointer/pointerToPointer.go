package main

import "fmt"

func main() {

	var ptr *int   // Declaring pointeer
	var pptr **int // Declaring pointer to pointer

	a := 5
	ptr = &a
	pptr = &ptr

	fmt.Println(a)
	fmt.Println(ptr)    // Address of a
	fmt.Println(pptr)   // Address of pointer
	fmt.Println(*ptr)   // value of a through pointer
	fmt.Println(**pptr) // valur of pointer through pointer which is a

}
