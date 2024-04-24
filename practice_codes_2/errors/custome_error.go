package main

import (
	"fmt"
	"log"
)

type wrongAge struct {
	age int
	msg string
}

func (e *wrongAge) Error() string {
	return fmt.Sprintf("\n %d: %s", e.age, e.msg)
}

func enterAge(age int) (string, error) {
	if age < 0 || age > 130 {
		return "", &wrongAge{age, "worng age value entered by user"}
	}
	return fmt.Sprintf("processing %d age value", age), nil
}

func main_custome() {
	var age int = 18
	msg, err := enterAge(age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)

	age = 178
	msg, err = enterAge(age)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
