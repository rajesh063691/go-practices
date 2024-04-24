package simpleinterest

import "fmt"

var name string = "Rajesh KUmar"

func init() {
	fmt.Println("Name", name)
	fmt.Println("Simple interest package initialized")
}

//Calculate calculates and returns the simple interest for a principal p, rate of interest r for time duration t years
func Calculate(p float64, r float64, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}
