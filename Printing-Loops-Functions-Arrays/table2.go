package main

import "fmt"

func main() {

	j := 1
	for i := 10; i < 20; i++ {
		fmt.Println(i, "*", j, "=", i*j)
		j = j + 1
	}
}
