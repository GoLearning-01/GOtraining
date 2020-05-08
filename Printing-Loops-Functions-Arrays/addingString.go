package main

import "fmt"

func main(){

	fmt.Println("Enter the first string")
	var firstString string
	fmt.Scanln(&firstString)

	fmt.Println("Enter the second string")
	var secondString string
	fmt.Scanln(&secondString)

	var stringAddition = firstString + secondString
	fmt.Println("The string addition is:", stringAddition )

}