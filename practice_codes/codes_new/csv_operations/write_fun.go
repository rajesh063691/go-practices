package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main_WriteFunc() {
	csvHeaders := []string{
		"FirstName", "LastName", "Occupation",
	}

	records := [][]string{
		{"John", "Doe", "gardener"},
		{"Lucy", "Smith", "teacher"},
		{"Brian", "Bethamy", "programmer"},
	}

	f, err := os.Create("userswrite.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	//write headers separately
	if err := w.Write(csvHeaders); err != nil {
		log.Fatalln("failed to write headers", err)
	}

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("failed to write in file", err)
		}
	}

}
