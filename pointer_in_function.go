package main

import "fmt"

type Address struct {
	City, Province, Country string
}

/*
	Pointers in Functions
	By default, function parameters in Go are passed by value,
	which means the data is copied and sent to the function.
	Therefore, changes made to the data inside the function 
	do not affect the original variable.
	This makes variables safe from accidental modification.

	However, sometimes we want a function to modify the original parameter.
	To achieve this, we can use pointers in function parameters.
	To declare a parameter as a pointer, use the * operator in the parameter type.
*/


// this won't change the Country, because the address is passed by value
// func ChangeCountryToIndonesia(address Address) {
// 	address.Country = "Indonesia"
// }

// this will change the Country, because the address is passed by reference (pointer)
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	
	ChangeCountryToIndonesia(&address)
	fmt.Println(address)
}