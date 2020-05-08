package main

import "fmt"

func swap(x, y string) (swapA string, swapB string) {
	swapA = y
	swapB = x
	return
}

func main() {

	var a string = "go"
	var b string = "lang"

	fmt.Println(a, b)

	fmt.Println(swap(a, b))

}
