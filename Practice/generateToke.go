package main

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func generateToken() string {
	b := make([]byte, 48)
	_, err := rand.Read(b)
	log.Println(err, "GenerateServiceToken??")
	token := base64.URLEncoding.EncodeToString(b)
	return token
}

func main_g() {
	// token := generateToken()
	// fmt.Println(token)

	// test := fmt.Sprintf("%s%s%s", "rajesh", "", "kumar")
	// fmt.Println(test)

	// val := 123.456
	// // fl, err := strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }
	// // fmt.Println(fl)
	// println(math.Round(val * 100))
	// fmt.Println(math.Round(val*100) / 100)

	//avg := getAverage(0.4555, 5)
	//fmt.Println(float64(int(0.4555 * 100)))
}

//	func getAverage(avg float64, count int) float64 {
//		if count == 0 {
//			return 0.0
//		}
//		average := avg / float64(count)
//		return float64(int(average*100)) / 100
//	}
// func getAverage(avg float64, count int) float64 {
// 	if count == 0 {
// 		return 0.0
// 	}
// 	average := avg / float64(count)
// 	f := float64(int(average * 100))
// 	return f
// }
