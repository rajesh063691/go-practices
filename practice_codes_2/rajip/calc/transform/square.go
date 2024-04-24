package transform

// return square of slices
func SquareSlice(s []int) (squareS []int) {

	if len(s) == 0 {
		return squareS
	}

	//make slice of length s
	squareS = make([]int, len(s))
	// for each element, save its sqaure
	for index, value := range s {
		squareS[index] = value * value
	}
	// returned transformed slice
	return squareS
}
