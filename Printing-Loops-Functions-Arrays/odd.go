package main

import (
	"fmt"
)

func main(){

	fmt.Println("Enter a number:")
	var i int
	fmt.Scanln(&i)

	if(i%2 != 0){
		fmt.Println("It is odd")
	}else{
		fmt.Println("It is even")
	}

}