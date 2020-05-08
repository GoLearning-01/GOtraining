package main

func main() {

	var i int = 0
	var oddSum int
	var evenSum int

	for i <= 10 {
		if i%2 == 0 {
			evenSum = evenSum + i
		} else {
			oddSum = oddSum + i
		}
		i++
	}

	println("Sum of even numbers from 1-10 =", evenSum)
	println("Sum of odd numbers from 1-10=", oddSum)

}
