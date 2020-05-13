package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	read, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File Content:", string(read))
}
