package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	myData := []byte("Hi there world!")

	err := ioutil.WriteFile("newFile.txt", myData, 0777)
	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("newFile.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
