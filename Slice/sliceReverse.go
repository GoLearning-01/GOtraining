package main

import "fmt"

func main() {

	var slice []int
	slice = append(slice, 3, 6, 9, 12, 15)

	var reverseSlice []int

	for i := len(slice) - 1; i >= 0; i-- {
		reverseSlice = append(reverseSlice, slice[i])
	}

	fmt.Println("Slice =", slice)
	fmt.Println("Reverse Slice =", reverseSlice)

}
