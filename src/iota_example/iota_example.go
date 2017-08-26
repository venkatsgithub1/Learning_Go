package main

import (
	"fmt"
)

const (
	messageToPrint = "%d %d\n"
	num1           = iota * 2
	num2
)

func main() {
	fmt.Printf(messageToPrint, num1, num2)
}
