package main

import "fmt"

func main() {
	// In the below 3 dots are used to for array to
	// automatically know the size.
	array1 := [...]int{0, 1, 1}
	fmt.Println(array1)

	// To know length of array len function can be used.
	fmt.Println(len(array1))

	// To make slice from an array.
	slice1 := array1[:] // This takes all elements of the array

	// if [1:] is used, data from index 1 to end of array.
	// if [:2] is used, data from index 0 to index 1, data at index 2 will not be included.

	fmt.Printf("slice1: %v\n", slice1)

	// appending data to a slice.

	slice1 = append(slice1, 2)
	fmt.Printf("slice after append operation: %v", slice1)
}
