package main

import "fmt"

func main() {
	// in const, we are allowed to not use the const, even though we have declared it
	// we can't change the value of constant
	// we need to spesify the initial value in the declaration
	const firstName string = "Adi"
	const lastName = "Wahyudi"

	// error, we can't change the value of constant
	// firstName = "Wahyudi"
	// lastName = "Adi"

	// we can declare multiple constant
	const (
		country = "Indonesia"
		hoby = "Coding"
	)
	fmt.Println(country)
	fmt.Println(hoby)
}
