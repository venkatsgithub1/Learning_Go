package main

import "fmt"

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

	m := map[string]person{
		person1.lastName: person1,
		person2.lastName: person2,
	}

	for k, v := range m {
		fmt.Println("key", k)
		fmt.Println(v.firstName)
		for idx, v2 := range v.favoriteIcecreamFlavors {
			fmt.Println("index", idx, "value:", v2)
		}
	}

}
