package main

import (
	"fmt"
)

const (
	message = "Hello, world!\n"
	afloat  = 3.14
)

var (
	message2     = "The value of pi is:%f\n"
	somevariable = 12
)

func main() {
	fmt.Printf("message: " + message)
	fmt.Printf(message2, afloat)
	somevariable += 13
	fmt.Printf("value is:%d", somevariable)
}
