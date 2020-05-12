//  EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.

package main

import "fmt"

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

func main() {

	gameSlice := []game{}

	entry1 := game{item{1, "god of war", 50}, "action adventure"}
	entry2 := game{item{2, "x-com 2", 30}, "strategy"}
	entry3 := game{item{3, "minecraft", 20}, "sandbox"}

	gameSlice = append(gameSlice, entry1, entry2, entry3)

	for _, v := range gameSlice {
		fmt.Println("")
		fmt.Println("ID:", v.item.id, "\nPrice:", v.item.price, "\nName:", v.item.name, "\nGenre:", v.genre)
	}
}
