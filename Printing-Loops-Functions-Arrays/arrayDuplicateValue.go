package main

import "fmt"

func main() {

	array1 := [10]string{"Atlanta", "Bentonville", "Charlotte", "Dallas", "Estes Park", "Forney", "Dallas", "Houston", "Indianapolis", "Jacksonville"}

	duplicate := ""

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array1); j++ {
			if i != j {
				if array1[i] == array1[j] {
					duplicate = array1[i]
				}
			}
		}
	}

	if duplicate != "" {
		fmt.Println("Duplicate value =", duplicate)
	} else {
		fmt.Println("No duplicate value!")
	}

}
