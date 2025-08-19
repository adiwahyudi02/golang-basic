package main

import "fmt"

func main() {
	name := "Adi"

	if name == "Adi" {
		fmt.Println("Hello Adi")
	} else if name == "Wahyudi" {
		fmt.Println("Hello Wahyudi")
	} else {
		fmt.Println("Hello, who are you?")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Name has correct length")
	}
}