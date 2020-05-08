package main

import "fmt"

func main() {

	fmt.Println("Hi")
	defer fmt.Println("World")
	fmt.Println("Hello")

}
