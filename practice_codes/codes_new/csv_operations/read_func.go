package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main_read() {
	f, err := os.Open("numbers.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(record)
		for value := range record {
			fmt.Printf("%s\n", record[value])
			//fmt.Printf("%d %s\n", i, value)
		}
	}
}
