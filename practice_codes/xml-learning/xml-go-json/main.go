package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Data struct {
	XMLName xml.Name `xml:"data" json:"-"`
	Person  []Person `xml:"person" json:"people"`
}

type Person struct {
	FirstName string   `xml:"firstname" json:"firstname"`
	LastName  string   `xml:"lastname" json:"lastname,omitempty"`
	Address   *Address `xml:"address" json:"address,omitempty"`
}

type Address struct {
	City  string `xml:"city" json:"city,omitempty"`
	State string `xml:"state" json:"state,omitempty"`
}

func main() {

	// xml to json conversion

	rawXmlData := `
		<data>
			<person>
				<firstname>Rajesh</firstname>
				<lastname>Kumar</lastname>
				<address>
					<city>Bangalore</city>
					<state>Karnataka</state>
				</address>
			</person>
			<person>
				<firstname>Puja</firstname>
				<lastname>Kumari</lastname>
				<address>
					<city>Kolkata</city>
					<state>West Bengal</state>
				</address>
			</person>
		</data>
	`
	var data Data

	err := xml.Unmarshal([]byte(rawXmlData), &data)
	if err != nil {
		fmt.Println(err)
	}

	jsonData, _ := json.MarshalIndent(data, "", "  ")

	fmt.Printf("JSON DATA: %s\n", string(jsonData))

	err = ioutil.WriteFile("data/rawJsonData.json", jsonData, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}

	// json to xml conversion

	ranJsonData := `
	{"people":[{"firstname":"Rajesh","lastname":"Kumar","address":{"city":"Bangalore","state":"Karnataka"}},{"firstname":"Puja","lastname":"Kumari","address":{"city":"Kolkata","state":"West Bengal"}}]}
`

	var newData Data

	err = json.Unmarshal([]byte(ranJsonData), &newData)
	if err != nil {
		fmt.Println(err)
		return
	}

	XmlByteData, err := xml.MarshalIndent(newData, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("XML DATA: %s\n", string(XmlByteData))
	//fmt.Println(string(XmlByteData))

	err = ioutil.WriteFile("./data/rawJsonData.xml", XmlByteData, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
}
