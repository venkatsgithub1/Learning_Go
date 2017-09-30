package main

import "fmt"

func main() {
	// make takes two or three parameters to create a slice
	// two parameter implementation.
	// make function creates a slice with a capacity of 100.
	slice1 := make([]float32, 100)

	slice1[0] = 2.
	slice1[1] = 4.
	slice1[2] = 6.
	slice1[3] = 8.

	fmt.Println(slice1)

	fmt.Printf("length of the slice: %d", len(slice1))
}
