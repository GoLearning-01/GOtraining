package main

import "fmt"

func main() {

	array1 := [9]int{1, 4, 6, 7, 3, 4, 6, 1, 7}
	slice1 := array1[0:9]

	for i := 0; i < len(slice1); i++ {
		for j := 0; j < len(slice1); j++ {
			if i != j {
				if slice1[i] == slice1[j] {
					if slice1[i] != 0 || slice1[j] != 0 {
						slice1[i] = 0
						slice1[j] = 0
					}
				}

			}
		}
	}

	for i, v := range slice1 {
		if v != 0 {
			fmt.Println("Unpaired element =", v, "; Index =", i)
		}
	}

}
