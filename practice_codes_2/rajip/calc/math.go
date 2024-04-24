package calc

import "errors"

//returns sum of integers
func Add(numbers ...int) (int, error) {
	sum := 0
	if len(numbers) < 2 {
		errMsg := "provide more than 2 numbers"
		return sum, errors.New(errMsg)
	} else {
		for _, num := range numbers {
			sum += num
		}
	}
	return sum, nil
}
