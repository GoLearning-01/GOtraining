package main

import "fmt"

func main() {
	var x string = "*"
	for i := 0; i < 5; i++ {
		fmt.Println(x)
		x = x + "*"
	}
}
