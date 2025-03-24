package calc_1

import "testing"

type TestingDataItems struct {
	inputs   []int
	result   int
	hasError bool
}{
	{[]int{1, 2, 3}, 6, false},
}

func TestAdd1(t *testing.T) {
	data := []TestingDataItems{
		{[]int{1, 2, 3}, 6, false},
		{[]int{99, 99}, 198, false},
		{[]int{1, 1, 5, 5, 6}, 18, false},
		{[]int{1, 2}, 4, true},
		{[]int{1}, 0, true},
	}

	for _, item := range data {
		count, err := Add(item.inputs...)
		if item.hasError {
			if err == nil {
				t.Errorf("Failed, expected=%v, got=%v", item.result, count)
			} else {
				t.Logf("Success, expected=%v, got=%v", item.result, count)
			}
		} else {
			if count != item.result {
				t.Errorf("Failed, expected=%v, got=%v", item.result, count)
			} else {
				t.Logf("Success, expected=%v, got=%v", item.result, count)
			}
		}
	}

}
