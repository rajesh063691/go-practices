package main

import "fmt"

func main_scanf() {
	name, age := "", 0
	fmt.Printf("Enter Your Name and Age \n")
	fmt.Scanf("%s  %d", &name, &age)
	fmt.Println("<===User Input Values====>")
	fmt.Printf("%s %d\n", name, age)
}
