package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Declare the arrays as slices
//
//   1. First run the following program as it is
//
//   2. Then change the array declarations to slice declarations
//
//   3. Observe whether anything changes or not (on the surface :)).
//
// EXPECTED OUTPUT
//  names    : []string ["Einstein" "Tesla" "Shepard"]
//  distances: []int [50 40 75 30 125]
//  data     : []uint8 [72 69 76 76 79]
//  ratios   : []float64 [3.14]
//  alives   : []bool [true false true false]
//  zero     : []uint8 []
// ---------------------------------------------------------

func main() {

	names := [3]string{"Einstein", "Tesla", "Shepard"}
	nameSlice := names[0:3]

	distances := [...]int{50, 40, 75, 30, 125}
	distanceSlice := distances[0:5]

	data := [5]byte{'H', 'E', 'L', 'L', 'O'}
	dataSlice := data[0:5]

	ratios := [1]float64{3.14145}
	ratioSlice := ratios[0]

	alives := [...]bool{true, false, true, false}
	aliveSlice := alives[0:4]

	zero := [0]byte{}
	zeroSlice := zero

	fmt.Printf("names    : %[1]T %[1]q\n", nameSlice)
	fmt.Printf("distances: %[1]T %[1]d\n", distanceSlice)
	fmt.Printf("data     : %[1]T %[1]d\n", dataSlice)
	fmt.Printf("ratios   : %[1]T %.2[1]f\n", ratioSlice)
	fmt.Printf("alives   : %[1]T %[1]t\n", aliveSlice)
	fmt.Printf("zero     : %[1]T %[1]d\n", zeroSlice)
}
