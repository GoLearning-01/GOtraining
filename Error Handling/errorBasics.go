package main

import (
	"errors"
	"fmt"
)

func calculateArea(radius int) (int, error) {
	if radius < 0 {
		return 0, errors.New("Provide Positive Number")
	}
	return radius * radius, nil
}

func main() {
	areaValue, err := calculateArea(-1)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error Encountered...")
		return
	}

	fmt.Println(areaValue)
}

/*

From the calling function main, thecalculateArea function is called with a negative parameter value.
Since the value passed is negative, we expect an “error” object from the function.
On function execution, it returns two values.
The first value represents the calculated area value and the second parameter represents the “error” object.
In the code above, we look for whether the function returns an “error” object or not,
if the “error” object returned is “nil”, the function executes further, otherwise,
we return the function after displaying the error message.

*/
