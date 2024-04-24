package main

import (
	"fmt"
)

// type Employee struct {
// 	firstName string
// 	lastName  string
// 	fullTime  bool
// 	salary    Salary
// }

// type Salary struct {
// 	basic     int
// 	allowance int
// 	bonus     Bonus
// }

// type Bonus struct {
// 	amount int
// 	yearly bool
// }

//type Employee2 org.Employee2

// type FullNameType func(string, string, string) string

// type Employee struct {
// 	FirstName, MiddleName, LastName string
// 	FullName                        FullNameType
// }

type LeaveType map[string]int

type Employee struct {
	firstName string
	lastName  string
	salary    int
	roll      int
	leaves    LeaveType
}

func main_prac() {
	//var emp1 Employee
	// emp.firstName = "Rajesh"
	// emp.lastName = "Kumar"
	// emp.salary = 100
	// emp.fullTime = true
	// emp := Employee{
	// 	firstName: "Rajesh",
	// 	lastName:  "Kumar",
	// 	salary:    100,
	// 	fullTime:  true,
	// }
	// assigning values without field names
	//emp := Employee{"Rajesh", "Kumar", 100, true}
	// defining anonymous struct type, it only takes struct and name is not required
	// emp := struct {
	// 	name   string
	// 	rollNo int
	// }{
	// 	name:   "Rajesh",
	// 	rollNo: 123,
	// }
	// fmt.Printf("%v", emp)
	// fmt.Printf("%T", emp)
	// fmt.Printf("%T", emp1)

	//creating pointer to struct
	// emp := &Employee{
	// 	firstName: "Rajesh",
	// 	lastName:  "Kumar",
	// 	salary:    1200,
	// 	fullTime:  true,
	// }

	//fmt.Println(emp)
	//fmt.Printf("%T", emp)
	//fmt.Println("firstName:", emp.firstName)
	// emp := Employee{
	// 	firstName: "Rajesh",
	// 	lastName:  "Kumar",
	// 	fullTime:  true,
	// 	salary: Salary{
	// 		basic:     1000,
	// 		allowance: 2000,
	// 		bonus: Bonus{
	// 			yearly: true,
	// 			amount: 1500,
	// 		},
	// 	},
	// }
	// fmt.Println(emp)
	// fmt.Println("emp firstName:", emp.firstName)
	// fmt.Println("emp basic salary:", emp.salary.basic)
	// fmt.Println("emp bonus amount:", emp.salary.bonus.amount)

	// emp2 := Employee2{
	// 	FirstName: "Rajesh",
	// 	LastName:  "Kumar",
	// 	//salary:    1000, //not able to set bcoz salary field is not exportable
	// }
	// fmt.Println(emp2)

	// emp := Employee{
	// 	FirstName:  "Rajesh",
	// 	MiddleName: "Kumar",
	// 	LastName:   "Yadav",
	// 	FullName: func(fsName, mdName, lsName string) string {
	// 		return fsName + " " + mdName + " " + lsName
	// 	},
	// }

	//fmt.Println(emp.FullName(emp.FirstName, emp.MiddleName, emp.LastName))
	leaveMaps := make(map[string]int)
	leaveMaps["Peresonal_leaves"] = 10
	leaveMaps["Medical_leaves"] = 12
	leaveMaps["Sick_leaves"] = 15

	emp1 := Employee{
		firstName: "Rajesh",
		lastName:  "Yadav",
		salary:    1200,
		roll:      1,
		leaves:    leaveMaps,
	}

	// emp2 := Employee{
	// 	firstName: "Rajesh",
	// 	lastName:  "Yadav",
	// 	salary:    1200,
	// 	roll:      1,
	// }

	//fmt.Println(emp1 == emp2)
	fmt.Println(emp1)
}
