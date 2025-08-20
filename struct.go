package main

import "fmt"

/*
	Struct

	A struct is a data template used to combine zero or more different data types into a single unit.
	Structs are usually used to represent data in the programs we create.
	Data in a struct is stored in fields.
	Simply put, a struct is a collection of fields.
*/

type Customer struct {
	Name, Address string
	Age int
}

func main() {

	/*
		Creating Struct Data

		A struct is a data template or prototype.
		A struct itself cannot be used directly.
		However, we can create data/objects from the struct that we have defined.
	*/

	var customer1 Customer
	customer1.Name = "Dadang"
	customer1.Address = "Bandung"
	customer1.Age = 20

	fmt.Println(customer1)
	fmt.Println(customer1.Name)
	fmt.Println(customer1.Address)
	fmt.Println(customer1.Age)

	// Struct Literal
	customer2 := Customer{
		Name: "Dudung",
		Address: "Cimahi",
		Age: 21,
	}
	fmt.Println(customer2)

	customer3 := Customer{"Diding", "Jakarta", 22}
	fmt.Println(customer3)
}