package main

import "fmt"

func main() {
	words := [...]string{"Alpha", "Bravo", "Charlie", "Delta"}
	simpsons := [5]string{"Homer", "Marge", "Bart", "Lisa", "Maggi"}
	fmt.Printf("%s\n", words[2])
	for _, simpson := range simpsons {
		fmt.Printf("%s\n", simpson)
	}
	fmt.Println("-----------------------")
	printArray(simpsons)
	fmt.Println("-----------------------")
	printArray(simpsons)
	fmt.Println("Array data is not changed")
	// Array passed to function is passed as a copy.
	// Use slice to modify data.
	// Using slice:
	simpsons2 := []string{"Homer", "Marge", "Bart", "Lisa", "Maggi"}
	fmt.Println("-----------------------")
	printSlice(simpsons2)
	fmt.Println("-----------------------")
	printSlice(simpsons2)
	fmt.Println("Slice data is changed")
	fmt.Println("-----------------------")
	printSlice(simpsons2[0:2]) // 0 and 1
	fmt.Println("-----------------------")
	printSlice(simpsons2[2:]) // 2 to end of slice
	// Using make keyword.
	cranes := make([]string, 3)
	cranes[0] = "Frasier"
	cranes[1] = "Martin"
	cranes[2] = "Niles"
	fmt.Println("-----------------------")
	printSlice(cranes)
	fmt.Println("-----------------------")
	currencies := make([]string, 0, 4) // starts with 0 items and can have
	// 4 items
	fmt.Printf("Length of slice: %d Capacity of slice: %d\n", len(currencies), cap(currencies))
	currencies = append(currencies, "USD")
	currencies = append(currencies, "GBP")
	currencies = append(currencies, "EUR")
	currencies = append(currencies, "JYN")
	fmt.Printf("Length of slice: %d Capacity of slice: %d\n", len(currencies), cap(currencies))
	// currencies[4]="RMB" fails.
	currencies = append(currencies, "RMB")
	fmt.Printf("Length of slice: %d Capacity of slice: %d\n", len(currencies), cap(currencies))
	// slice capacity increases with initial capacity.
	printSlice(currencies)
	// slice copy
	currencies_copy := make([]string, len(currencies))
	copy(currencies_copy, currencies)
	fmt.Println("-----------------------")
	printSlice(currencies_copy)
	currencies_copy[2] = "AUD"
	fmt.Println("-----------------------orig----------------")
	printSlice(currencies)
	fmt.Println("-----------------------copy----------------")
	printSlice(currencies_copy)
	fmt.Println("-----------------------")
	fmt.Println("currencies_copy is a new slice")
	fmt.Println("-----------------------")
	fmt.Println("reference copy")
	currencies_ref := currencies[0:3]
	printSlice(currencies_ref)
	currencies_ref[0] = "SGD"
	fmt.Println("-----------------------orig----------------")
	printSlice(currencies)
	fmt.Println("-----------------------REF----------------")
	printSlice(currencies_ref)
	fmt.Println("change in ref slice changes original slice")
}

func printSlice(sliceToPrint []string) {
	for _, str := range sliceToPrint {
		fmt.Printf("%s\n", str)
	}
	//sliceToPrint[2]="Burt"
}

func printArray(arrayToPrint [5]string) {
	for _, str := range arrayToPrint {
		fmt.Printf("%s\n", str)
	}
	arrayToPrint[2] = "Burt"
}
