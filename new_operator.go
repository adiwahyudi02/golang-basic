package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	/*
		Operator new

		Previously, we used the & operator to create a pointer to an existing variable.
		Go also provides the built-in function new to create a pointer.
		However, new only returns a pointer to zeroed (empty) memory, 
		meaning the data has no initial value.
	*/


	// var address1 *Address = &Address{}
	var address1 *Address = new(Address)
	var address2 *Address = address1

	address2.Country = "Indonesia"
	fmt.Println(address1) // The adderess1 will change to "Indonesia"
	fmt.Println(address2) // The adderess2 will change to "Indonesia"
}