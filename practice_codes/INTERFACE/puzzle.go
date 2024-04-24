package main

import "fmt"

type puzzle struct {
	title string
	price int64
}

func (g *puzzle) print() {
	fmt.Printf("%-10s:%d \n", g.title, g.price)
}
