package main

import (
	"fmt"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	nameSlice := strings.Split(namesA, ", ")

	for i := 0; i < len(namesB); i++ {
		for j := 0; j < len(nameSlice); j++ {
			if namesB[i] == nameSlice[j] {
				namesB[i] = ""
			}
		}
	}

	count := 0
	for i := 0; i < len(namesB); i++ {
		if namesB[i] != "" {
			count = 1
			fmt.Println("The slices are not equal!")
		}
	}

	if count != 1 {
		fmt.Println("The slices are equal!")
	}

}
