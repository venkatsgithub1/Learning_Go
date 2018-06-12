package main

import "fmt"

func main() {
	x := make([][]string, 0)
	x = append(x, []string{"James", "Bond", "Shanken, not stirred"})
	x = append(x, []string{"Miss", "Moneypenny", "Helloooooo, James."})
	fmt.Println(x)
	for i, v := range x {
		fmt.Println(i, v)
	}

	// or

	xs1 := []string{"James", "Bond", "Shanken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	xxs := [][]string{xs1, xs2}

	for _, v := range xxs {
		for j, v1 := range v {
			fmt.Println(j, v1)
		}
	}
}
