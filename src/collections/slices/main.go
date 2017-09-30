package main

import "fmt"

func main() {
	// Making a slice from scratch.
	// No need to provide three dots or number for size in square brackets.
	slice1 := []float32{2., 4., 6., 8.}
	fmt.Println(slice1)

	// len function can be used to know the length of the slice.
	fmt.Printf("length of the slice: %d", len(slice1))
}
