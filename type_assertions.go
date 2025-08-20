package main

import "fmt"

/*
	Type Assertions

	Type assertions allow us to convert a variable from an interface type
	to a more specific type.
	This feature is often used when working with empty interface (`any`) data,
	to extract the actual underlying value and type.
*/

func random() any {
	return true
}

func main() {
	// var result any = random()
	// var message string = result.(string)
	// fmt.Println(message)

	// panic: interface conversion: interface {} is string, not int
	// var result2 any = random()
	// var message2 int = result2.(int)
	// fmt.Println(message2)

	/*
		Type Assertions Using Switch

		Using a type assertion incorrectly can cause a panic in your application.
		If the panic is not recovered, the program will terminate.
		To make type assertions safer, it is recommended to use a type switch expression,
		which allows handling different types without causing a panic.
	*/

	result := random()
	switch message2 := result.(type) {
	case string:
		fmt.Println("String: ", message2)
	case int:
		fmt.Println("Integer: ", message2)
	default:
		fmt.Println("Unknown type")
	}
}