package main

import ("fmt"
		"net/http"
		"io/ioutil"
		"html/template"
		"encoding/xml"
		"sync"
)

var wg sync.WaitGroup

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

// we don't want to pass in the News struct to this function, mainly cause you don't want to read/write to a obj running concurrently
func newsRoutine(c chan News, Location string) {
	// This will yield this value to the wait() command when finished
	defer wg.Done()

	var n News

	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	// Unmarshal maps an XML element to a Name by recording the element name.
	// Unmarshal values, from the body response, and send it to memory address of SitemapIndex instance s
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()

	// Add values to the channel ( remember, channel is kinda like a ... queue in python ?)
	c <- n
}

// w = writer
// r = reading through pointer to http.Request value
func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	// ################################ Begin - Taken from gotut15 ##################################
	// Create instance of struct SitemapIndex ( kinda like initalizing a class )
	var s SitemapIndex

	// NOTE: Create a map where the key is a string, and the values are of type struct NewsMap
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// Unmarshall values, from the body response, and send it to memory address of SitemapIndex instance s
	xml.Unmarshal(bytes, &s)
	news_map := make(map[string]NewsMap)
	resp.Body.Close()

	// TODO: Lookup this make command
	// Buffer just needs to be bigger than the amount you're using
	queue := make(chan News, 30)

	// Grab data from each news link in sitemap
	fmt.Printf("Here %s some %s", "are", "variables")
	for _, Location := range s.Locations {
		wg.Add(1)
		go newsRoutine(queue, Location)
	}

	// Wait for goroutine to finish
	wg.Wait()

	// close channel when finished
	close(queue)
	// ################################ END - Taken from gotut15 ##################################

	// Iterate over the channel
	for elem := range queue {
		// idx means index
		for idx, _ := range elem.Titles {
			news_map[elem.Titles[idx]] = NewsMap{elem.Keywords[idx], elem.Locations[idx]}
		}
	}

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
