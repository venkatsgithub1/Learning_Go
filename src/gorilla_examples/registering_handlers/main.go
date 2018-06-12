package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func users(w http.ResponseWriter, r *http.Request) {
	htmlString := fmt.Sprint("<html><div>Al</div><br><div>Bob</div></html>")
	fmt.Fprintf(w, htmlString)
}

func posts(w http.ResponseWriter, r *http.Request) {
	postsString := fmt.Sprint("<html>" +
		"<div>2018 Formula One Canadian GP is on June 10th.</div><br>" +
		"</html>")
	fmt.Fprintf(w, postsString)
}

func home(w http.ResponseWriter, r *http.Request) {
	homePageString := "<h1>Welcome to demo blog</h1>"
	fmt.Fprintf(w, homePageString)
}

func main() {
	// creating a new gorilla mux router.
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/users", users)
	r.HandleFunc("/posts", posts)

	http.Handle("/", r)

	http.ListenAndServe(":8000", nil)
}
