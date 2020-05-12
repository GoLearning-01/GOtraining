package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
}

type Teacher struct {
	firstName string
	lastName  string
	subject   string
}

func main() {

	student1 := Student{"Summit", "Khatiwada", 24}
	student2 := Student{firstName: "Mason", lastName: "Pandey", age: 5}

	teacher1 := Teacher{"Apil", "Pandey", "Mathematics"}

	fmt.Println("")
	fmt.Println(student1)
	fmt.Println("")
	fmt.Println(student2)
	fmt.Println("")
	fmt.Println(teacher1.lastName, "->", teacher1.subject)

}
