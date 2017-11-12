package main

// If we didn't have the washtion post website available, we could uncomment this below
/*
var washPostXML = []byte(`
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
   <sitemap>
      <loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
   </sitemap>
   <sitemap>
      <loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
   </sitemap>
</sitemapindex>`)
*/

import ("fmt"
		"net/http"
		"io/ioutil"
		"encoding/xml")

type SitemapIndex struct {
	// MUST capitalize values, otherwise they won't get exported, and you won't be able to use them
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

// Array example: [5]type == array
// Slice example: []type == slice ( In python, this would be a list? Not sure yet )

// value receiver ( This is kinda like overriding __str__ in python or something like that. Give a string representation of a Struct )
func (L Location) String() string {
	// Sprintf formats according to a format specifier and returns the resulting string.
	return fmt.Sprintf(L.Loc)
}

func main() {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Create instance of struct SitemapIndex ( kinda like initalizing a class )
	var s SitemapIndex

	// Unmarshall values, from the body response, and send it to memory address of SitemapIndex instance s
	xml.Unmarshal(bytes, &s)

	// At this point, we have an array of sorts, not string
	fmt.Println(s.Locations)
}
