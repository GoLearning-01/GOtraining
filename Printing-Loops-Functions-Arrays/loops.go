package main

import (
	"fmt"
)

func main(){

	fmt.Println("Enter the first number: ")
	var firstNum int
	fmt.Scanln(&firstNum)

	fmt.Println("Enter the second number: ")
	var secondNum int
	fmt.Scanln(&secondNum)

	fmt.Println("Enter the thirs number: ")
	var thirdNum int
	fmt.Scanln(&thirdNum)

	if (firstNum > secondNum && firstNum > thirdNum){
		fmt.Println(firstNum, "is the biggest")
	}

	if (secondNum > firstNum && secondNum > thirdNum){
		fmt.Println(secondNum, "is the biggest")
	}

	if (thirdNum > firstNum && thirdNum > secondNum){
		fmt.Println(thirdNum, "is the biggest")
	}

}
