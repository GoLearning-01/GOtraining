package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number

	phoneBook := make(map[string]string)
	phoneBook["Summit"] = "215-917-9380"
	phoneBook["Aashish"] = "580-695-6920"
	phoneBook["Ankit"] = "569-092-0191"
	phoneBook["Sophie"] = "214-2412-2672"

	fmt.Println("")
	fmt.Println("Phonebook:")
	for i, v := range phoneBook {
		fmt.Println(i, "->", v)
	}

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable

	product := map[int]string{
		001: "Available",
		002: "Available",
		003: "Out of Stock",
		004: "Available",
	}

	fmt.Println("")
	fmt.Println("Stock:")
	for i, v := range product {
		fmt.Println("Product ID:", i, "is", v)
	}

	// #3
	// Key        : Last name
	// Element    : Phone numbers

	phoneBook2 := make(map[string]string)
	phoneBook2["White"] = "320-099-0929"
	phoneBook2["Brown"] = "320-499-0959"
	phoneBook2["Gray"] = "312-099-0322"
	phoneBook2["Green"] = "320-129-0929"

	fmt.Println("")
	fmt.Println("Phonebook2:")
	for i, v := range phoneBook2 {
		fmt.Println(i, "->", v)
	}

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity

	basket := make(map[string]string)
	basket["Rodd"] = "6 Bananas"
	basket["Vince"] = "7 Apples"
	basket["Tommy"] = "10 Shoes"
	basket["CJ"] = "15 Shirts"

	fmt.Println("")
	for i, v := range basket {
		fmt.Println("Mr.", i, "bought", v)
	}

}
