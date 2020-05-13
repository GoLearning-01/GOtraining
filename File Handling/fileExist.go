package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		fmt.Println("The file does not exist")
		return
	}
	fmt.Println("The file exists!")
	//perform operations on the file/file contents
}
