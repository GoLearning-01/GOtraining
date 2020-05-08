package main

import "fmt"

func main() {

	array := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	even := 0
	odd := 0

	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 {
			even = even + 1
		} else {
			odd = odd + 1
		}
	}
	fmt.Println("Total even =", even)
	fmt.Println("Total odd =", odd)
}
