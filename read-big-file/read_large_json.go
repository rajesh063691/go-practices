package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Item struct {
	Data map[string]interface{}
}

func main_L() {

	file, err := os.Open("single.json")
	if err != nil {
		fmt.Println("error opening the file")
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("error getting file info")
		return
	}

	fmt.Println(fileInfo.Size())

	// create a new json decoder
	decoder := json.NewDecoder(file)

	// Expect the JSON data to be an array
	// Read the opening bracket

	// read the token only when you have array of json objects, else simply do not check the token condition
	var token json.Token
	if token, err = decoder.Token(); err != nil {
		fmt.Println("Error reading JSON", err)
		return
	}

	fmt.Printf("TOKEN START  =%v \n", token)

	// Iterate over the array elements

	count := 0

	for decoder.More() {
		count++
		var item Item
		//var item map[string]interface{}

		if err := decoder.Decode(&item.Data); err != nil {
			fmt.Println("Error decoding JSONson", err)
			return
		}

		fmt.Printf("DATA :: %+v\n", item)
		fmt.Println("<====================>")
		if url, ok := item.Data["type"].(string); ok {
			fmt.Printf("URL :: %s\n", url)
		}

		v := reflect.ValueOf(item.Data["actor"])
		fmt.Println(v)
		fmt.Println(v.Kind())
		if v.Kind() == reflect.Map {
			fmt.Printf("ACTOR :: %+v\n", v)
			avatarURL := v.MapIndex(reflect.ValueOf("avatar_url"))
			fmt.Printf("AVATAR URL :: %s\n", avatarURL)
		}
		//fmt.Printf("URL :: %s\n", item.Data["avatar_url"])
		fmt.Println("<====================>")
		fmt.Println("<====================>")
		if count > 10 {
			break
		}
	}

	// Read the closing bracket
	if token, err = decoder.Token(); err != nil {
		fmt.Println("Error reading JSON", err)
		return
	}

	fmt.Printf("TOKEN END = %v \n", token)

}
