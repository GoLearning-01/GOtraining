package main

import "fmt"

func main() {

	var newMap = make(map[string]int)

	newMap["Alice"] = 2761781718
	newMap["Bob"] = 8276511717

	for name, number := range newMap {
		if name == "Bob" {
			fmt.Println(name, number)
		}
	}
}
