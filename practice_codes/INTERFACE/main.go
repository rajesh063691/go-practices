package main

import "fmt"

func main_1() {
	var (
		game1   = game{title: "game-1", price: 200}
		game2   = game{title: "game-2", price: 300}
		book1   = book{title: "book-1", price: 500}
		book2   = book{title: "book-2", price: 600}
		puzzle1 = puzzle{title: "puzzle-1", price: 700}
		puzzle2 = puzzle{title: "puzzle-2", price: 800}
	)
	var store list

	store = append(store, &game1, &game2, &book1, &book2, &puzzle1, &puzzle2)
	store.print()

	// we can compare interface values as well
	fmt.Println(store[0] == &game1)
	fmt.Println(store[3] == &book2)
}
