package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type person struct {
	firstName               string
	lastName                string
	favoriteIcecreamFlavors []string
}

func main() {
	person1 := person{
		firstName:               "James",
		lastName:                "Bond",
		favoriteIcecreamFlavors: []string{"chocolate", "butterscotch", "vanilla"},
	}

	person2 := person{
		firstName:               "Miss",
		lastName:                "Moneypenny",
		favoriteIcecreamFlavors: []string{"strawberry", "pineapple", "peachfruit"},
	}

	spew.Dump(person1)
	spew.Dump(person2)

	printSlice := func(person person) {
		fmt.Println(person.firstName)
		fmt.Println(person.lastName)
		for i, v := range person.favoriteIcecreamFlavors {
			fmt.Println("index", i, "value", v)
		}
	}

	printSlice(person1)
	printSlice(person2)

}
