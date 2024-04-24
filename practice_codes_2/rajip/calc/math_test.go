package calc

import "testing"

type TestDataItems struct {
	inputs   []int
	result   int
	hasError bool
}

func TestAdd(t *testing.T) {
	dataSet := []TestDataItems{
		{[]int{1, 2, 3}, 6, false},
		{[]int{99, 99}, 198, false},
		{[]int{1, 1, 5, 5, 6}, 18, false},
		{[]int{1}, 0, true},
	}
	for _, item := range dataSet {
		result, err := Add(item.inputs...)
		if item.hasError {
			if err == nil {
				t.Errorf("Faild, expected=%v, got=%v", item.result, result)
			} else {
				t.Logf("Success, expected=%v, got=%v", item.result, err)
			}

		} else {
			if result != item.result {
				t.Errorf("Faild, expected=%v, got=%v", item.result, result)
			} else {
				t.Logf("Success, expected=%v, got=%v", item.result, result)
			}
		}
	}
}

func TestAdd_1(t *testing.T) {
	sum, _ := Add(1, 2, 3)
	if sum != 6 {
		t.Errorf("Add(1, 2, 3) Failed, expected=%v, got=%v", 6, sum)
	} else {
		t.Logf("Add(1, 2, 3) Success, expected=%v, got=%v", 6, sum)
	}

	sum, _ = Add(1)
	if sum != 0 {
		t.Errorf("Add(1) Failed, expected=%v, got=%v", 0, sum)
	} else {
		t.Logf("Add(1) Success, expected=%v, got=%v", 0, sum)
	}
}
