package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main_0() {
	// websites := [...]string{"webcode.me", "zetcode.com", "freebsd.org", "netbsd.org"}

	// re := regexp.MustCompile(`(\\w+)\\.(\\w+)`)

	// for _, website := range websites {

	// 	parts := re.FindStringSubmatch(website)

	// 	for i, _ := range parts {
	// 		fmt.Println(parts[i])
	// 	}

	// 	fmt.Println("---------------------")
	// }

	// re := regexp.MustCompile(`\d{4}.\d{2}.\d{2}`)

	// fmt.Printf("Pattern: %v\n", re.String()) // print pattern

	// fmt.Println(re.MatchString(str1)) // true

	// submatchall := re.FindAllString(str1, -1)
	// println(submatchall[0])
	// // for _, element := range submatchall {
	// // 	fmt.Println(element)
	// // }

	str1 := "mcmcmccm-2022.09.14-b8d20d8-dc207be-2.20220914.1661418144"
	//str1 := "dash-chart-master-bb49e30-391a513-2.02345354.1661508976"
	fmt.Println("Checking relase and IsProd for string:", str1)
	release, isProd := FindRelease(str1)
	fmt.Printf("Release=[%v] and isProd=[%v]\n", release, isProd)

}

func FindRelease(releaseFormat string) (release string, isProduction bool) {
	log.Println("Inside FindRelease")

	// checking release for prod
	//re := regexp.MustCompile(`\d{4}.\d{2}.\d{2}`)
	re := regexp.MustCompile(`\d{4}.(0[1-9]|1[0-2]).(0[1-9]|[12][0-9]|3[01])`)
	if re.MatchString(releaseFormat) {
		allMatches := re.FindAllString(releaseFormat, -1)
		return allMatches[0], true
	}
	// find release for non-prod
	releaseSlice := strings.Split(releaseFormat, "-")

	return releaseSlice[len(releaseSlice)-1], false
}
func getReleaseName1(relFormat, revision string) (relName string) {
	if relFormat == "" {
		return
	}
	if strings.ToLower(relFormat) == "yyyymmdd" {
		if strings.Contains(revision, ".") {
			revSlice := strings.Split(revision, ".")
			//valid release name
			if len(revSlice) > 1 && len(revSlice[1]) == 8 {
				for i, val := range strings.Split(revSlice[1], "") {
					relName = relName + val
					if i == 3 || i == 5 {
						relName = relName + "."
					}
				}
			}
		}

	}
	return
}
