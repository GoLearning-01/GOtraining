package main

import "fmt"

func main() {

	a := [5]int{1, 2, 3, 4, 5}

	var b [len(a)]int

	j := 0
	for i := len(a) - 1; i >= 0; i-- {
		b[j] = a[i]
		j++
	}

	for i := 0; i < len(a); i++ {
		a[i] = b[i]
	}

	fmt.Println(a)
}
