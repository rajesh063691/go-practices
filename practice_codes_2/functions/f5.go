// Go program to illustrate the function
// as a field in Go structure
package main

import "fmt"

func CalculateFinalSal(mat, pay int) int {
	return mat * pay
}

// Finalsalary of function type
type Finalsalary func(int, int) int

// Creating structure
type Author struct {
	name      string
	language  string
	Marticles int
	Pay       int
	// Function as a field,
	//ou can set it by defining type or you can directly set using function name which has the actual definition
	//1st way of declaration
	Salary Finalsalary

	//2nd way of declaration
	//Salary func(int, int) int

	//both trype of declaration works
}

func (au *Author) getAuthorDetails() {
	name := au.name
	language := au.language
	materials := au.Marticles
	pay := au.Pay
	Salary := au.Salary(materials, pay)

	// Display values
	fmt.Println("Author's Name: ", name)
	fmt.Println("Language: ", language)
	fmt.Println("Total number of articles published in May: ", materials)
	fmt.Println("Per article pay: ", pay)
	fmt.Println("Total salary: ", Salary)
}

func setAuthor() *Author {
	return &Author{
		name:      "Sonia",
		language:  "Java",
		Marticles: 120,
		Pay:       500,
		Salary:    CalculateFinalSal,
	}
}

// Main method
func main_f5() {
	// Initializing the author fields
	author := setAuthor()

	fmt.Printf("Author Object=%+v", author)

	//fmt.Printf("get Salary:%v \n", author.Salary(author.Marticles, author.Pay))

	//another way of accessing boith gives same result
	author.getAuthorDetails()

}
