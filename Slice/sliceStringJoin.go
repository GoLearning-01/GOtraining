package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	names := []string{
		"Einstein", "Shepard",
		"Tesla"}

	// -----------------------------------
	books := []string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}

	sort.Strings(books)

	// -----------------------------------
	// // this time, do not change the nums array to a slice
	num := [...]int{5, 1, 7, 3, 8, 2, 6, 9}

	// // use the slicing expression to change the nums array to a slice below
	nums := num[0:8]
	sort.Ints(nums)

	// -----------------------------------
	// Here: Use the strings.Join function to join the names
	//       (see the expected output)
	fmt.Println(strings.Join(names, " and "))

	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}
