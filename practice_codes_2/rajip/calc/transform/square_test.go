package transform

import (
	"reflect"
	"testing"
)

func TestSquareSlice(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	expextedResult := []int{1, 4, 9, 16, 25}

	result := SquareSlice(testSlice)
	if reflect.DeepEqual(expextedResult, result) {
		t.Logf("PASSED!")
	} else {
		t.Errorf("Failed, expected=%v, got=%v", expextedResult, result)
	}
}

func TestSquareSlice_1(t *testing.T) {
	testSlice := []int{}
	expextedResult := []int{}

	result := SquareSlice(testSlice)
	if reflect.DeepEqual(expextedResult, result) {
		t.Logf("PASSED!")
	} else {
		t.Errorf("Failed, expected=%v, got=%v", expextedResult, result)
	}
}
