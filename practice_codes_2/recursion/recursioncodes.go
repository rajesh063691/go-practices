package recursion

import "fmt"

var i = 1

func TestRecursion(val int) {
	fmt.Println(val)
	//base case
	if val <= 1 {
		return
	} else {
		//recursive case
		TestRecursion(val - 1)
	}

}

func Factorial(n int) int {
	//base case
	if n == 0 {
		return 1
	}
	//recursive case
	result := Factorial(n - 1)
	return n * result
}

func Fibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}

}

func PrintNumbers(n int) {
	if n <= 50 {
		fmt.Println(n)
		PrintNumbers(n + 1)
	}
}

func PrintNumbers2(start, end int) int {
	//base case
	if start == end {
		return end
	} else {
		//recursive case
		return start + PrintNumbers2(start+1, end)
	}
}

func PrintFibanocci(preNum, fibNum, endNum int) int {
	nextNum := 0
	//base case
	if i == endNum {
		return 0
	} else {
		nextNum = preNum + fibNum
		preNum = fibNum
		fibNum = nextNum
		i++
		fmt.Println(fibNum)
		PrintFibanocci(preNum, fibNum, endNum)
	}
	return 0
}
func PrintArray(arr []int, i int) {
	//base case
	if i > len(arr)-1 {
		return
	} else {
		fmt.Println(arr[i])
		PrintArray(arr, i+1)
	}
}

func CountDigit(n int) int {
	//base case
	if n <= 0 {
		return 0
	} else {
		n = n / 10
		//recursive case
		return 1 + CountDigit(n)
	}
}

func SumOfDigit(n int) int {
	//base case
	if n <= 0 {
		return 0
	} else {
		r := n % 10
		n = n / 10
		//recursive case
		return r + SumOfDigit(n)
	}
}

func FindMaxInArray(arr []int, i, n, max int) int {
	if i < n {
		if max < arr[i] {
			max = arr[i]
			i++
			return FindMaxInArray(arr, i, n-i, max)
		}
	}
	return 0
}
