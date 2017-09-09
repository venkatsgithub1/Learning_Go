package main

import (
	"fmt"
)

func main() {
	printer("Hello, World!")
	fmt.Println("---------------------------------------")
	printer_multiple_params("Hello, ", "World!", 5)
	fmt.Println("---------------------------------------")
	data, _ := fmt.Printf("%s\n", "counthis")
	fmt.Println("---------------------------------------")
	fmt.Println(data)
	fmt.Println("---------------------------------------")
	printer_with_return("Hello World!")
	fmt.Println("\n---------------------------------------")
	appendedMessage, err := printer_with_multi_return("Hello, World!")
	if err == nil {
		fmt.Printf("%q\n", appendedMessage)
	}
	fmt.Println("---------------------------------------")
	printer_with_defer("Hello, World!")
	fmt.Println("---------------------------------------")
	printer_with_named_return_val("Named return method")
	returnedData, err := printer_with_named_return_multiple_vals("Named multiple returns")
	fmt.Println("---------------------------------------")
	fmt.Printf("%s %v\n", returnedData, err)
	fmt.Println("---------------------------------------")
	printer_dot_dot_dot("Hello, ", "World!")
	fmt.Println("---------------------------------------")
	printer_dot_dot_dot_mul_params("%x\n", "Hello Go", "in Hex format")
}

func printer_dot_dot_dot_mul_params(format string, msgs ...string) {
	for _, msg := range msgs {
		fmt.Printf(format, msg)
	}
}

func printer_dot_dot_dot(msgs ...string) {
	for _, msg := range msgs {
		fmt.Printf("%s\n", msg)
	}
}

func printer_with_named_return_multiple_vals(msg string) (returnedMsg string, e error) {
	_, e = fmt.Printf("%s\n", msg)
	returnedMsg = "Returned Message"
	return
}

func printer_with_named_return_val(msg string) (e error) {
	_, e = fmt.Printf("%s\n", msg)
	return
}

func printer_with_defer(msg string) error {
	defer fmt.Printf("\n")
	defer fmt.Printf("Will be printed after msg\n")
	_, err := fmt.Printf("%s", msg)
	return err
}

func printer(msg string) {
	fmt.Printf("%s\n", msg)
}

//func printer_multiple_params(msg string, msg2 string) { or
func printer_multiple_params(msg, msg2 string, repeat int) {
	for repeat > 0 {
		fmt.Printf("%s", msg)
		fmt.Printf("%s\n", msg2)
		repeat--
	}
}

func printer_with_return(msg string) error {
	_, err := fmt.Printf("%s", msg)
	return err
}

func printer_with_multi_return(msg string) (string, error) {
	msg += "\n"
	_, err := fmt.Printf("%s", msg)
	return msg, err
}
