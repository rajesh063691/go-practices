package exerc01

import (
	"fmt"

	excelize "github.com/xuri/excelize/v2"
)

func CreateNewSheet() {
	f := excelize.NewFile()
	sheetNames := []string{"Transcript", "Transcript_Rajesh"}
	f.SetSheetName("Sheet1", sheetNames[0])
	f.SetSheetName("Sheet2", sheetNames[1])
	f.SetActiveSheet(1)

	data := [][]interface{}{
		{"Student Exam Score"},
		{"Type: Mid Term Exam", nil, nil, nil, "Core Curriculam", nil, nil, "Science"},
		{"Number", "ID", "Name", "Class", "Language Arts", "Mathematics", "History", "Chemistry", "Biology", "Physics", "Total"},
		{1, 10001, "Student A", "Class 1", 93, 80, 89, 86, 57, 77},
		{2, 10002, "Student B", "Class 2", 93, 80, 89, 86, 57, 77},
		{3, 10003, "Student C", "Class 3", 93, 80, 89, 86, 57, 77},
		{4, 10004, "Student D", "Class 4", 93, 80, 89, 86, 57, 77},
		{5, 10005, "Student E", "Class 5", 93, 80, 89, 86, 57, 77},
		{6, 10006, "Student F", "Class 6", 93, 80, 89, 86, 57, 77},
	}

	for i, row := range data {
		startCell, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, sheetName := range sheetNames {
			if err := f.SetSheetRow(sheetName, startCell, &row); err != nil {
				fmt.Println(err)
				return
			}
		}

	}
	if err := f.SaveAs("Book2.xlsx"); err != nil {
		fmt.Println(err)
		return
	}
}
