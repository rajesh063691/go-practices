package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main_WriteAll() {
	records := [][]string{
		{"first_name", "last_name", "occupation"},
		{"John", "Doe", "gardener"},
		{"Lucy", "Smith", "teacher"},
		{"Brian", "Bethamy", "programmer"},
	}

	f, err := os.Create("usersWriteAll.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	err = w.WriteAll(records)
	if err != nil {
		log.Fatalln("failed to writeAll in file", err)
	}
}
