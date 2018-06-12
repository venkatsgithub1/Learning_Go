package main

import (
	"fmt"
	"strings"
)

func main() {
	wordCount := func(s string) map[string]int {
		counts := map[string]int{}
		for _, word := range strings.Fields(s) {
			counts[word]++
		}
		return counts
	}

	str := "some data to be checked for count"
	fmt.Println(wordCount(str))
}
