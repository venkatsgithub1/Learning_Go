package main

import (
	"fmt"
	"net/http"
)

// HandlerFunc has logic to show things for the path, client has requested.
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	// setting content type.
	w.Header().Set("Content-Type", "text/html")
	// checking path.
	path := r.URL.Path
	if path == "/" {
		fmt.Fprintf(w, "<h1>Welcome to path /</h1>")
	} else if path == "/contact" {
		fmt.Fprintf(w, "John Doe can be reached at <a href='mailto:johndoe@none.com></a>.")
	} else {
		fmt.Fprintf(w, "<h1>Snap!</h1><p>Cannot find the page, you're looking for.</p>")
	}
}

func main() {
	http.HandleFunc("/", HandlerFunc)
	http.ListenAndServe(":3000", nil)
}
