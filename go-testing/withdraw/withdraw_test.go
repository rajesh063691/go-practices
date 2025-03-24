package withdraw

import "testing"

func TestWithdraw(t *testing.T) {
	var tests = []struct {
		name           string
		currentBalance int
		withdrawAmount int
		newBalance     int
	}{
		{"Withdraw 10 from 100", 100, 10, 90},
		{"Withdraw 20 from 100", 100, 20, 80},
		{"Withdraw 30 from 100", 100, 30, 70},
		//{"Withdraw 40 from 100", 100, 40, 600},
		{"Withdraw 150 from 100", 150, 150, 150},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			newBalance := Withdraw(test.currentBalance, test.withdrawAmount)
			if newBalance != test.newBalance {
				t.Errorf("Expected %d but got %d", test.newBalance, newBalance)
			}
		})
	}
}
