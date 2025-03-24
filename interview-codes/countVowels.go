package main

import "fmt"

func main_vowel() {
	mp := make(map[rune]int)
	str := "rajeshrajeshaa"
	countVowels(mp, str)

	for k, v := range mp {
		fmt.Printf("%s=%d\n", string(k), v)
	}
}

func countVowels(mp map[rune]int, str string) {
	for _, r := range str {
		if k, ok := mp[r]; ok {
			mp[r] = k + 1
		} else {
			mp[r] = 1
		}
	}
}
