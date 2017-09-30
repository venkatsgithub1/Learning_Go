package main

import "fmt"

func main() {
	// map can be created using make function
	// usage: inside make function provide
	// 1. map keyword.
	// 2. key type in square brackets.
	// 3. value type after square brackets.
	map1 := make(map[int]string)

	fmt.Println("empty map: ", map1)

	map1[5] = "Sebastian Vettel"
	map1[7] = "Kimi Raikkonen"
	map1[44] = "Lewis Hamilton"
	map1[77] = "Valtteri Bottas"

	fmt.Println("map: ", map1)

	// if key is not there in map and we try to access ?
	fmt.Println(map1[1])
	// An empty string, which is zero value for string will be printed out.
}
