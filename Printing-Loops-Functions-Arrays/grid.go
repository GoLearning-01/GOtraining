package main

import "fmt"

func main() {

	var grid [2][10]string

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = "-"
		}
	}

	for i := 0; i < len(grid); i++ {
		var output string = ""
		for j := 0; j < len(grid[i]); j++ {
			output = output + grid[i][j]
		}
		fmt.Println(output)
	}

}
