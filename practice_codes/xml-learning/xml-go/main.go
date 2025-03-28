package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

var washPostXML = []byte(`
<sitemapindex>
	<sitemap>
		<loc>https://test-1.xml</loc>
	</sitemap>
	<sitemap>
		<loc>https://test-1.xml</loc>
	</sitemap>
	<sitemap>
		<loc>https://test-1.xml</loc>
	</sitemap>
</sitemapindex>`)

type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")

	//bytes, _ := ioutil.ReadAll(resp.Body)
	bytes := washPostXML
	defer resp.Body.Close()

	var s SiteMapIndex
	err := xml.Unmarshal(bytes, &s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", s)
}
