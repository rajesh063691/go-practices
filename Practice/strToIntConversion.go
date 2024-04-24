package main

import (
	"fmt"
	"time"
)

func main_str() {
	// previousString := "1"
	// previousDays, err := strconv.Atoi(previousString)
	// if err != nil {
	// 	fmt.Printf("err=%s", err)
	// 	return
	// }
	// fmt.Printf("Previoudays=%d \n", previousDays)
	//fmt.Println(getCategory(0.88))

	//location, _ := time.LoadLocation("Asia/Calcutta")

	startedAt := "2024-02-26T05:40:15Z"
	completedAt := "2024-02-26T05:40:15.32Z"
	t1, _ := time.Parse(time.RFC3339, startedAt)
	//t1, _ = time.Parse("2006-01-02T15:04:05.99Z", t1.In(location).UTC().Format("2006-01-02T15:04:05.99Z"))

	t2, _ := time.Parse("2006-01-02T15:04:05.99Z", completedAt)
	//t2, _ = time.Parse("2006-01-02T15:04:05.99Z", t2.In(location).UTC().Format("2006-01-02T15:04:05.99Z"))

	diff := t2.UnixMilli() - t1.UnixMilli()

	// fmt.Println(t2)
	// fmt.Println(t2.UTC())
	// fmt.Println(t1.UTC().UnixMilli())
	// fmt.Println(t2.Unix(), t2.UnixMilli())
	// fmt.Println(t1.Unix(), t1.UnixMilli())
	// fmt.Println(t1)
	// fmt.Println(t2)
	fmt.Println(int64(diff))
	t := t2.Sub(t1)
	epoc_time := time.Now().AddDate(0, 0, -360).UnixNano()
	fmt.Printf("%s \n", t)
	fmt.Print(epoc_time, "\n")
}

// func getCategory(avg float64) (category string) {
// 	// avg is average for 7 days
// 	// once per week means 1/7 = 0.14
// 	// once per week month 1/30 = 0.03
// 	avgPerDay := float64(int((avg/7)*100)) / 100
// 	if avgPerDay > 1 {
// 		return "Elite"
// 	} else if avgPerDay >= 0.14 && avgPerDay <= 1 {
// 		return "High"
// 	} else if avgPerDay >= 0.03 && avgPerDay < 0.14 {
// 		return "Medium"
// 	}
// 	return "Low"
// }

// day>1 or week >7   - Elite
// month>30  - Elite
// 60
