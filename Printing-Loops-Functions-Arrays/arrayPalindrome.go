package main

import "fmt"

// I copied this function form medium.com
func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {

	array := [7]string{"cat", "madam", "liril", "dog", "rotor", "matter", "racecar"}

	fmt.Println("Palindromes in array:")

	for i := 0; i < len(array); i++ {
		if array[i] == reverse(array[i]) {
			fmt.Println(array[i])
		}
	}

}
