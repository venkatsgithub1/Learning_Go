package main

import "fmt"

func main() {
	/*var counter int
	counter = 0
	for counter < 10 {
		fmt.Printf("Hello, World!\n")
		counter += 1
	}*/

	/*for counter := 0; counter < 10; counter++ {
		fmt.Printf("Hello, World!\n")
	}*/
	i := 0
	var stop bool

	for !stop {
		fmt.Printf("Hello, World!\n")
		i++
		if i == 10 {
			stop = true
		}
	}
}
