package main

import (
	"fmt"
	"runtime"
)

func main() {
	num1 := 1

	if num1 == 1 {
		fmt.Println("One")
	} else if num1 == 2 {
		fmt.Println("Two")
	} else {
		fmt.Println("Not One or Two")
	}

	if num2 := 3; num2 == 4 {
		fmt.Println("Four")
	} else if num2 == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("Not Three or Four")
	}

	// Below example can be looked at in golang online.
	// if a case is matched, the body of swtich
	// breaks automatically.
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux OS")
	case "darwin":
		fmt.Println("OS X")
	case "windows":
		fmt.Println("Windows OS")
	default:
		fmt.Println("OS is not Linux or OS X or Windows")
	}

	num3 := 10

	// Another switch usage example.
	switch {
	case num3 == 1:
		fmt.Println("One")
	case num3 > 5:
		fmt.Println("Number is greater than 5")
	}
}
