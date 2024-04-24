package main

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectDetails struct {
	ID              primitive.ObjectID
	Name            string
	ApplicationName string
}

var projectMap map[string]ProjectDetails

func main_4() {
	// prj := ProjectDetails{
	// 	Name:            "Default",
	// 	ApplicationName: "test-1",
	// }
	mp := make(map[string]ProjectDetails, 0)
	prj := new(ProjectDetails)
	prj.Name = "default"
	prj.ApplicationName = "app1"
	mp[prj.Name] = *prj

	fmt.Printf("Initial Map=%v\n", mp)

	prj.Name = "default"
	prj.ApplicationName = "app2"
	mp[prj.Name] = *prj

	fmt.Printf("Initial Map=%v\n", mp)

	prj.Name = "default"
	prj.ApplicationName = "app3"
	mp[prj.Name] = *prj

	fmt.Printf("Initial Map=%v\n", mp)

	for key, val := range mp {
		fmt.Printf("Key=%s\n", key)
		fmt.Printf("Val=%v\n", val)
	}

	version := strings.Split("2.20220921.1662468688", ".")
	fmt.Printf("len=%d&&&val=%s", len(version), version[2])

	str1 := "631ebf2b01f615e3b0cf248a"
	str2 := "2.20220928.1662973138"

	first_5_digit := str1[0:5]
	last_5_digit := str2[(len(str2) - 5):]
	fmt.Println()
	fmt.Println("First 5=", first_5_digit)
	fmt.Println("Last 5=", last_5_digit)
}
