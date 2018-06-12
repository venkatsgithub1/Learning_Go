package main

import "fmt"

func first() {
	fmt.Println("1")
}

func second() {
	fmt.Println("2")
}

func main() {
	// function second will be called at the end of the function.
	defer second()
	first()
	fmt.Println("function first executed")
	fmt.Println("function second will execute now")
}
