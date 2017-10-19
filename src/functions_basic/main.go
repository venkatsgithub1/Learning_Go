package main

import "fmt"

func main() {
	// pass by value
	dataToFunc := "Go"
	messageByVal(dataToFunc)
	fmt.Println(dataToFunc)

	messageByRef(&dataToFunc)
	fmt.Println(dataToFunc)

	variadicFunc("Hi", "Go")
}

// pass by value
func messageByVal(data string) {
	fmt.Printf("Hi %s\n", data)
	data = "There"
}

// pass by reference
func messageByRef(data *string) {
	fmt.Printf("Hi %s\n", *data)
	*data = "There"
}

// variadic function
func variadicFunc(data ...string) {
	for _, dataItem := range data {
		fmt.Println(dataItem)
	}
}
