package withdraw

import "fmt"

func Withdraw(currentBalance, withdrawAmount int) int {
	if withdrawAmount > currentBalance {
		fmt.Println("Insufficient funds.")
		return currentBalance
	}

	newBalance := currentBalance - withdrawAmount
	fmt.Printf("Withdrawal of %d successful. New balance: %d\n", withdrawAmount, newBalance)
	return newBalance
}
