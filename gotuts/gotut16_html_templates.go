package main

import ("fmt"
		"net/http"
		"html/template"
)

type NewsAggPage struct {
	Title string
	News string
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
	p := NewsAggPage{Title: "Amazing News Aggregator", News: "Some news"}

	// create template now. t = template, _ = errors (panic etc)
	t, _ := template.ParseFiles("basictemplating.html")

	t.Execute(w,p)
}

func main() {
	// Create handlers to deal with events after hitting routes in web app
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/agg/", newsAggHandler)

	// Create actual server for application
	http.ListenAndServe(":8000", nil)
}
