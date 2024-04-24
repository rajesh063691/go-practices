package main

import "fmt"

type book struct {
	title string
	price int64
}

func (g *book) print() {
	fmt.Printf("%-10s:%d \n", g.title, g.price)
}
