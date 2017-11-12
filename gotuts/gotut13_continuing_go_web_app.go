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

func main() {

	// Create instance of struct SitemapIndex ( kinda like initalizing a class )
	var s SitemapIndex
	var n News

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Unmarshall values, from the body response, and send it to memory address of SitemapIndex instance s
	xml.Unmarshal(bytes, &s)

	fmt.Printf("Here %s some %s", "are", "variables")
	for _, Location := range s.Locations {
		resp, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		// Unmarshall values, from the body response, and send it to memory address of SitemapIndex instance s
		xml.Unmarshal(bytes, &n)
	}
}


// Next tutorial, we'll build a map then display to page
