package main

import (
	"fmt"
)

type Student struct {
	rollno int
	name   string
}

func main() {
	// instance of structure
	s := Student{101, "Jack"}

	//pointer to th student struct
	ps := &s
	fmt.Println(ps)
	fmt.Println("hi")

	//Accessing struct fields via pointer
	fmt.Println((*ps).name)
	fmt.Println(ps.name)
	(*ps).rollno = 322 // place *ps in () while using in struct!
	fmt.Println(ps)
}
