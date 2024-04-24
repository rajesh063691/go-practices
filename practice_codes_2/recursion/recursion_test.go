package recursion

import "testing"

func TestFactorial(t *testing.T) {
	//t.isParallel
	val := Factorial(4)
	if val != 120 {
		t.Errorf("Error in test code")
	}
}

func TestPrintFibanocci(t *testing.T) {
	//t.isParallel
	val := PrintFibanocci(0, 1, 5)
	_ = val

	// if val != 120 {
	// 	t.Errorf("Error in test code")
	// }
}
func TestCountDigit(t *testing.T) {
	//t.isParallel
	val := CountDigit(123)
	//_ = val

	if val != 3 {
		t.Errorf("Error in test code")
	}
}
func TestSumOfDigit(t *testing.T) {
	//t.isParallel
	val := SumOfDigit(12345)
	//_ = val

	if val != 15 {
		t.Errorf("Error in test code")
	}
}

func TestFindMaxInArray(t *testing.T) {
	//t.isParallel
	val := FindMaxInArray([]int{5, 6, 8, 9, 10, 11}, 0, 6, 0)
	//_ = val

	if val != 15 {
		t.Errorf("Error in test code")
	}
}
