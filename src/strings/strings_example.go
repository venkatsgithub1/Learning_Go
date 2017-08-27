package main

import (
	"fmt"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"
	fmt.Printf("%s\n", atoz[:12])
	fmt.Printf("%s\n", atoz[13:19])
	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}
	// No need of index here.
	for _, r := range atoz {
		fmt.Printf("%c\n", r)
	}
	// optional second keyword in for removed.
	for i := range atoz {
		fmt.Printf("%d\n", i)
	}
	// Length of the string.
	fmt.Printf("%d is the length of atoz\n", len(atoz))
    // Below is an example for literal strings.
	atoz2 := `the quick brown fox jumps over the lazy dog\n`
	fmt.Printf("%s\n", atoz2)
}
