package main

import "fmt"

func main_6() {

	uniqueVal := []string{"uid-1", "uid-2", "uid-3", "uid-1"}
	apps := make(map[string]map[string]int)
	for _, st := range uniqueVal {

		val, ok := apps["project"][st]
		fmt.Printf("%v %v\n", val, ok)
		if ok {
			fmt.Printf("Value of project map :: %v", val)
		}
		if apps["project"] == nil {
			apps["project"] = make(map[string]int)

		}
		apps["project"][st] = 123
		fmt.Println(apps)

	}

	fmt.Printf("Map=%v", apps)

}
