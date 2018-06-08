package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to the home page</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "John Doe can be reached at this <a href='mailto:johndoe@none.com'>Email Address</a>.")
}

func main() {
	// we create a new router variable using gorilla mux package.
	r := mux.NewRouter()

	// pass handler function to the router.
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	// pass router variable to the http ListenAndServe method.
	http.ListenAndServe(":3000", r)
}
