// create and writeString

package main

import (
	"fmt"
	"os"
)

func main() {

	// Creating a file
	file, err := os.Create("firstFile.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Writing string on a file
	write, err := file.WriteString("Enter text here!")
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	fmt.Println(write, "bytes written sucessfully!")
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
