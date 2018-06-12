package main

import "github.com/davecgh/go-spew/spew"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"None": 0,
		},
		favDrinks: []string{
			"Coke",
			"Sprite",
		},
	}

	spew.Dump(s)
}
