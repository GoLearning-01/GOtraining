// divisible by 4 is leap year
// century divisible by 100 is not
// century divisible by 400 is

package main

import "fmt"

func main() {

	fmt.Println("Enter year:")
	var year int
	fmt.Scanln(&year)

	if year%100 == 0 {
		if year%400 == 0 {
			fmt.Println("It is a leap year")
		} else {
			fmt.Println("It is not a leap year")
		}
	} else {
		if year%4 == 0 {
			fmt.Println("It is a leap year")
		} else {
			fmt.Println("It is not a leap year")
		}
	}
}
