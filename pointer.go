package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	/*
		Pass by Value

		By default, all variables in Go are passed by value, not by reference.
		This means that when we pass a variable to a function, method, or another variable,
		a copy of the value is actually sent, not the original variable.
	*/

	// address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	// address2 := address1 // copy the value

	// address2.City = "Cimahi"

	// fmt.Println(address1) // The adderess1 won't change
	// fmt.Println(address2) // The adderess2 will change became "Cimahi"


	/*
		Pointer

		A pointer is the ability to create a reference to a variable's memory location
		without duplicating the actual data.
		Simply put, using pointers allows us to achieve pass-by-reference behavior,
		so that changes made via the pointer affect the original variable.
	*/

	/*
		Operator &
		To create a pointer variable that holds the memory address of another variable,
		we use the & operator followed by the variable name.
	*/

	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer

	address2.City = "Cimahi"

	fmt.Println(address1) // The adderess1 will change became "Cimahi"
	fmt.Println(address2) // The adderess2 will change became "Cimahi"
}