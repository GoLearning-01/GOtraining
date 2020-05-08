package main

import "fmt"

func main() {

	a1 := [6]int{1, 2, 3, 4, 5, 6}
	a2 := [3]int{7, 8, 9}

	const l = len(a1) + len(a2)

	var a3 [l]int

	for i := 0; i < len(a1); i++ {
		a3[i] = a1[i]
	}

	count := len(a1)
	for i := 0; i < len(a2); i++ {
		a3[count] = a2[i]
		count++
	}

	fmt.Println(a3)

}
