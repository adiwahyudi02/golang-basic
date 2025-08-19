package main

import "fmt"

func main() {
	name := "Adi"

	switch name {
	case "Adi":
		fmt.Println("Hello Adi")
	case "Wahyudi":
		fmt.Println("Hello Wahyudi")
	default:
		fmt.Println("Hello, who are you?")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name is too long")
	default:
		fmt.Println("Name has correct length")
	}


	// switch without condition
	length := len(name)

	switch {
	case length > 5:
		fmt.Println("Name is too long")
	default:
		fmt.Println("Name has correct length")
	}
}