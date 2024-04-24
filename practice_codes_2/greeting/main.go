package main

import (
	"fmt"

	"github.com/logrusorgru/aurora"
)

func main() {
	greetMessageEmpty := hello("")
	fmt.Println(aurora.Yellow(greetMessageEmpty))

	greetMessageRajesh := hello("Rajesh")
	fmt.Println(aurora.Yellow(greetMessageRajesh))
}

//returns user as string
func hello(user string) string {
	if len(user) == 0 {
		return "Hello Dude!"
	} else {
		return fmt.Sprintf("Hello %v!", user)
	}

}
