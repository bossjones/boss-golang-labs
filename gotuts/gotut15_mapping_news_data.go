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
	// sitemap>loc: location tag under the sitemap tag
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword string
	Location string
}

func main() {

	// Create instance of struct SitemapIndex ( kinda like initalizing a class )
	var s SitemapIndex
	var n News

	// NOTE: Create a map where the key is a string, and the values are of type struct NewsMap
	news_map := make(map[string]NewsMap)
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Unmarshall values, from the body response, and send it to memory address of SitemapIndex instance s
	xml.Unmarshal(bytes, &s)

	// Grab data from each news link in sitemap
	fmt.Printf("Here %s some %s", "are", "variables")
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		// Unmarshal maps an XML element to a Name by recording the element name.
		// Unmarshal values, from the body response, and send it to memory address of SitemapIndex instance s
		xml.Unmarshal(bytes, &n)

		// idx means index
		for idx, _ := range n.Titles {
			news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
		}
	}

	// Print out data now
	for idx, data := range news_map {
		fmt.Println("\n\n\n", idx)
		fmt.Println("\n", data.Keyword)
		fmt.Println("\n", data.Location)
	}
}


// Next tutorial, we'll build a map then display to page
