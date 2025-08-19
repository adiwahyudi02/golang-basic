package main

import "fmt"

func main() {
	// we can create our own type, it's called type declaration
	// it's an alias to another type
	type NoKTP string
	var noKTP NoKTP = "1234567890"

	var stringKTP string = "0987654321"

	// convert string to NoKTP
	var noKTP2 = NoKTP(stringKTP)

	fmt.Println(noKTP)
	fmt.Println(noKTP2)
}