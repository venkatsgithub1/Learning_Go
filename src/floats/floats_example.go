package main

import (
	"fmt"
)

func main() {
	var pi float32 = 3.14
	fmt.Printf("value:%f\n", pi)
	theoconst := 1.732
	fmt.Printf("theodorus constant:%f\n", theoconst)
	eumasconst := float64(0.57721)
	fmt.Printf("euler-mascheroni constant:%f\n", eumasconst)
	var pythagorasconst float64 = 1.414
	fmt.Printf("pythagaros constant:%f\n", pythagorasconst)
	fmt.Printf("pythagaros constant with two decimal points:%.2f\n", pythagorasconst)
}
