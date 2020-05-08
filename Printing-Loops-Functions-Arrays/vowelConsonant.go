package main

import "fmt"

func main() {

	fmt.Println("Enter an alphabet")
	var alphabet string
	fmt.Scanln(&alphabet)

	if len(alphabet) == 1 {

		if alphabet == "a" || alphabet == "e" || alphabet == "i" || alphabet == "o" || alphabet == "u" || alphabet == "A" || alphabet == "E" || alphabet == "I" || alphabet == "O" || alphabet == "U" {
			fmt.Println("It is a vowel")
		} else {
			fmt.Println("It is a consonant")
		}

	} else {
		fmt.Println("Input Invalid!")
	}

}
