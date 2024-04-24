package main

import (
	"fmt"
	"strings"
)

func main_1() {
	relFormat := "release-YYYY.MM.DD"
	revision := "2.20220901.1660940618"
	relName := getReleaseName(relFormat, revision)
	fmt.Printf("Release Name:=%s\n", relName)
}

func getReleaseName(relFormat, revision string) (relName string) {
	if relFormat == "" {
		return
	}
	relFormat = strings.ToLower(relFormat)
	if strings.Contains(revision, ".") {
		revSlice := strings.Split(revision, ".")
		//valid release name
		year, month, day := "", "", ""
		if len(revSlice) > 1 && len(revSlice[1]) == 8 {
			for i, val := range strings.Split(revSlice[1], "") {
				if i <= 3 {
					year = year + val
				}
				if i > 3 && i <= 5 {
					month = month + val
				}
				if i > 5 {
					day = day + val
				}
			}

			replacer := strings.NewReplacer("yyyy", year, "mm", month, "dd", day)
			relName = replacer.Replace(relFormat)
		}
	}

	return
}
