package main

import (
	"encoding/csv"
	"os"
)

func main() {
	records := [][]string{
		{"First Name", "Middle Nme", "Last Name"},
		{"Rajesh", "Kumar", "Yadav"},
		{"Pooja", "Kumari", "Yadav"},
		{"Prisha", "Kumari"},
	}

	//file, err := os.Create("./record.csv")
	file, err := os.Create("./record1.csv")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	w := csv.NewWriter(file)
	// if you want to print one by one
	// defer w.Flush()
	// for _, record := range records {
	// 	if err := w.Write(record); err != nil {
	// 		log.Fatalln("error writing record to file: ", err)
	// 	}
	// }

	// if you want to print all together
	err = w.WriteAll(records)
}
