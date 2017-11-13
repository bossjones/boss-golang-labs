package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"html/template"
		"encoding/xml"
)

type NewsAggPage struct {
	Title string
	// Map where key is a string, and value is struct NewsMap
	News map[string]NewsMap
}

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

// w = writer
// r = reading through pointer to http.Request value
func indexHandler(w http.ResponseWriter, r *http.Request) {

	// Format print
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

// w = writer
// r = reading through pointer to http.Request value
func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	// ################################ Begin - Taken from gotut15 ##################################
	// Create instance of struct SitemapIndex ( kinda like initalizing a class )
	var s SitemapIndex
	var n News

	// NOTE: Create a map where the key is a string, and the values are of type struct NewsMap
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	news_map := make(map[string]NewsMap)
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
	// ################################ END - Taken from gotut15 ##################################

	p := NewsAggPage{Title: "Amazing News Aggregator", News: news_map}

	// create template now. t = template, _ = errors (panic etc)
	t, _ := template.ParseFiles("newsaggtemplate.html")

	fmt.Println(t.Execute(w,p))
}

func main() {
	// Create handlers to deal with events after hitting routes in web app
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)

	// Create actual server for application
	http.ListenAndServe(":8000", nil)
}
