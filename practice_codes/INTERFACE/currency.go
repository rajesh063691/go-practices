package main

import "fmt"

type Currency float64
type Stringer interface {
	String() string
}

func (c *Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(*c))
}

func main() {
	// Currency implements  main.Stringer interface
	var c1 Currency = 1.5
	fmt.Printf("c's value=%v, Type=%T \n", c1.String(), c1.String()) //explicitly call the Currency String method
	fmt.Printf("c's value=%v, Type=%T  \n", c1, c1)                  // implicitly call's Currency String method by fmt package

	var c = new(Currency)
	*c = 1.5
	fmt.Printf("c1's value=%v  Type=%T\n", c.String(), c.String())
	fmt.Printf("c1's value=%v   Type=%T\n", c, c)

	var mainStringer Stringer
	mainStringer = c
	fmt.Printf("mainStringer's value=%v and type=%T \n", mainStringer, mainStringer)

	var fmtStringer fmt.Stringer
	fmtStringer = c
	fmt.Printf("fmtStringer's value=%v and type=%T \n", fmtStringer, fmtStringer)
}
