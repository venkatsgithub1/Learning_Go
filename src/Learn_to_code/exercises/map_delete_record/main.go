package main

import "fmt"

func main() {
	x := map[string][]string{"bond_james": []string{"Shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"}}

	fmt.Println(x)

	x["fleming_ian"] = []string{"snakes", "cigars", "espionage"}

	delete(x, "no_dr")

	for key, value := range x {
		fmt.Println("key:", key)
		for i, v := range value {
			fmt.Println("index:", i, "value:", v)
		}
	}
}
