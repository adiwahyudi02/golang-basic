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

/*
	Struct Method

	A struct is a data type like any other and can be used as a parameter for functions.
	However, if we want to add behavior to structs, we can define methods for them,
	making it seem like the struct “has” functions.
	A method is essentially a function with a receiver, which associates it with a specific struct.
*/
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
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


	customer1.sayHello("Dodong")
	customer2.sayHello("Dodong")
	customer3.sayHello("Dodong")
}