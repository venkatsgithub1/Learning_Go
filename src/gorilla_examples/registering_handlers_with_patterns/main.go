package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func test(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["a"])
	fmt.Fprintf(w, "Id provided in query: %s", vars["a"])
}

func homePage(w http.ResponseWriter, r *http.Request) {
	data := fmt.Sprint("Welcome to demo blog.")
	fmt.Fprintf(w, data)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homePage)
	r.HandleFunc("/test/{a:[0-9]+}", test)

	http.Handle("/", r)

	http.ListenAndServe(":8000", nil)
}
