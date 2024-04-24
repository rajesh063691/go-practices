package main

import "fmt"

type inputError struct {
	missingFields string
	message       string
}

// this is the actual error interface
func (e *inputError) Error() string {
	return e.message
}

func (e *inputError) getMissingField() string {
	return e.missingFields
}

func main_3() {
	ie := &inputError{
		message:       "my name is Rajesh",
		missingFields: "none",
	}
	err := validate("Rajesh", "none")
	if err != nil {
		if err, ok := err.(*inputError); ok {
			fmt.Println(err.message)
			fmt.Printf("Missing field is: %s\n", err.getMissingField())
			fmt.Println(err.Error())
		}
	}
	fmt.Printf("Response: %+v \n", *(ie))
}

func validate(name string, gender string) error {
	if name == "" {
		return &inputError{
			message:       "Name is mandatory",
			missingFields: "Name",
		}
	}

	if gender == "" {
		return &inputError{
			message:       "Gender is mandatory",
			missingFields: "Gender",
		}
	}
	return nil
}
