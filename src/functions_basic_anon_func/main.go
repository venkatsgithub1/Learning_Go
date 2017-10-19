package main

import "fmt"

func main() {
	addItems := func(items ...int) (countOfItems int, result int) {
		for _, item := range items {
			result += item
		}
		countOfItems = len(items)
		return
	}

	count, sum := addItems(10, 20, 30, 40, 50)
	fmt.Printf("%d numbers are added and their sum is: %d\n", count, sum)
}
