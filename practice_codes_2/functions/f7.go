package main

import "fmt"

type output func(string) string

func hello(name string) string {

	return fmt.Sprintf("hello %s", name)
}

func main_f7() {
	var f output
	fmt.Println(f)
	f = hello
	// it will give packagename.function name if custom type defined using type keyword else would giv efunction signature
	fmt.Printf("name=%T \n", f)
	fmt.Println(f("Peter"))
}
