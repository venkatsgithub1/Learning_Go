package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for as while
	fmt.Println("----------for as while-------------")
	j := 0
	for {
		j++
		fmt.Println(j)
		if j > 5 {
			break
		}
	}

	fmt.Println("---Looping with slices---")
	s := []string{"One", "Two", "Three"}
	for idx, v := range s {
		fmt.Printf("Index: %d Value: %s\n", idx, v)
	}

	fmt.Println("---Looping with maps---")
	map1 := make(map[int]string)
	map1[5] = "Seb"
	map1[7] = "Kimi"
	map1[44] = "Lewis"
	map1[77] = "Valtteri"

	for key, val := range map1 {
		fmt.Printf("Key: %d Value: %s\n", key, val)
	}
}
