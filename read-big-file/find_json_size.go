package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// data := map[string]interface{}{
	// 	"id":            "123456",
	// 	"name":          "Rajesh Kumar",
	// 	"qualification": "Software Engineer",
	// }

	data := `{"name":"Alice","age":30,"address":"123 Main St"}`

	byteData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error marshalling the data :: %v\n", err)
		return
	}

	fmt.Printf("TOTAL SIZE :: %d \n", len(byteData))
}
