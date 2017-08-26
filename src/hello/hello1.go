package main

import (
	"fmt"
)

func main() {
	//var message string
	//message := "Hello, world!\n"
	message := "Hello, world!\n"
	fmt.Printf("message: " + message)
	afloat := 3.14
	message2 := "The value of pi is:%f"
	fmt.Printf(message2, afloat)
}
