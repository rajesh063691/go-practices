package main

import "fmt"

type game struct {
	title string
	price int64
}

func (g *game) print() {
	fmt.Printf("%-10s:%d \n", g.title, g.price)
}
