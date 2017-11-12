package main

import ("fmt"
        "net/http")

// w = writer
// r = reading through pointer to http.Request value
func index_hander(w http.ResponseWriter, r *http.Request) {

	// Format print
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

// w = writer
// r = reading through pointer to http.Request value
func about_handler(w http.ResponseWriter, r *http.Request) {

	// Format print
	fmt.Fprintf(w, "Expert web design by bossjones")
}

func main() {
	// Create handlers to deal with events after hitting routes in web app
	http.HandleFunc("/", index_hander)
	http.HandleFunc("/about/", about_handler)

	// Create actual server for application
	http.ListenAndServe(":8000", nil)
}
