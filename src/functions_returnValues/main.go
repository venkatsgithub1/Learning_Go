package main

import "fmt"

func main() {
	fmt.Println(addIntegers(10, 20, 30, 40, 50))
	count, sum := addIntegersAndReturnSumAndItemCount(100, 90, 80, 70, 60)
	fmt.Printf("Number of integers: %d and their sum: %d\n", count, sum)
	count2, sum2 := namedParams(1000, 900, 800, 700, 600)
	fmt.Printf("Number of integers: %d and their sum: %d\n", count2, sum2)
}

func addIntegers(items ...int) int {
	result := 0
	for _, item := range items {
		result += item
	}

	return result
}

// Big Name of a function not really suggested, but this is for
// the sake of clarity.
func addIntegersAndReturnSumAndItemCount(items ...int) (int, int) {
	result := 0
	for _, item := range items {
		result += item
	}

	return len(items), result
}

func namedParams(items ...int) (countOfItems int, result int) {
	for _, item := range items {
		result += item
	}

	countOfItems = len(items)

	// Notice that no return variables aren't specified.
	// Just return statement will do.
	return
}
