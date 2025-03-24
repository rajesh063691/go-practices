package main

import (
	"fmt"
	"math/rand"
)

func main_a() {
	// var a = 10
	// func() {
	// 	fmt.Println("Hello")
	// }()

	// fmt.Print(a)

	// var anFunc = func() {
	// 	fmt.Println("Hello")
	// }

	// anFunc()

	// var addFunc = func(val1, val2 int) int {
	// 	return val1 + val2
	// 	//fmt.Printf("Sum of %d and %d :: %d \n", val1, val2, val1+val2)
	// }

	// fmt.Println(addFunc(10, 20))

	// 	var anFunc = func(st1, st2 string) string {
	// 		return st1 + st2 + "Geeks"
	// 	}

	// 	CFG(anFunc)
	// }

	// func CFG(i func(p, q string) string) {
	// 	fmt.Printf(i("Geeks", "For"))

	fmt.Println(Random())

}

var creatures = []string{"shark", "jellyfish", "squid", "octopus", "dolphin"}

func Random() string {
	i := rand.Intn(len(creatures))
	return creatures[i]
}
