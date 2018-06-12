package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func main() {
	r := mux.NewRouter()

	sr := r.Host("www.example.com").Subrouter()

	sr.HandleFunc("/", homePage)

	http.Handle("/", sr)

	http.ListenAndServe(":8000", nil)

}
