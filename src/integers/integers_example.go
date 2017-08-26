package main

import (
	"fmt"
)

func main() {
	nineplussix := 15
	fmt.Printf("value: %d\n", nineplussix)
	twelveplusthirteen := uint(25)
	fmt.Printf("twenty five:%d\n", twelveplusthirteen)
	oneplustwo := uint8(3)
	fmt.Printf("three:%d\n", oneplustwo)
	oneplus32767 := uint32(32768)
	fmt.Printf("thirty two thousand seven hundred and sixty seven:%d\n", oneplus32767)
	oneplus1048575 := uint64(1048576)
	fmt.Printf("1 million forty eight thousand five hundred and seventy six:%d\n", oneplus1048575)
}
