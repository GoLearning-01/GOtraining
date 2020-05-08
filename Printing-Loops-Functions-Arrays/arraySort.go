package main

import (
	"fmt"
	"sort"
)

func main() {

	array := []string{"philip", "charles", "harry", "elizabeth", "diana", "meghan", "william"}

	sort.Strings(array)

	fmt.Println(array)
}
