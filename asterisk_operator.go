package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // pointer
	address2.City = "Cimahi"
	fmt.Println(address1) // The adderess1 will change became "Cimahi"
	fmt.Println(address2) // The adderess2 will change became "Cimahi"


	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// fmt.Println(address1) // The adderess1 won't change, still "Cimahi"
	// fmt.Println(address2) // The adderess2 will change "Jakarta"


	/*
		Operator *

		When we change a pointer variable itself, only that pointer changes.
		Other variables pointing to the same data remain unaffected.
		If we want to modify the actual data being pointed to, 
		we can use the * operator to dereference the pointer.
	*/

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) // The adderess1 will change to "Jakarta"
	fmt.Println(address2) // The adderess2 will change to "Jakarta"
}