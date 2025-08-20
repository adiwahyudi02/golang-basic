package main

import "fmt"

/*
	Interface

	An interface is an abstract data type; it does not have a direct implementation.
	An interface contains definitions of methods.
	Interfaces are usually used as a contract, specifying behavior that other types must implement.
*/

type HasName interface {
	GetName() string
}

func SayHello (value HasName) {
	fmt.Println("Hello", value.GetName())
}


/*
	Implementing Interfaces

	Any data type that matches the method signatures of an interface 
	is automatically considered to implement that interface.
	This means we do not need to explicitly declare that a type implements an interface.
	This is different from some other programming languages, 
	where you must explicitly state that a type implements a specific interface.
*/

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {

	person := Person{Name: "Adi"}
	SayHello(person)

	animal := Animal{Name: "Cat"}
	SayHello(animal)
}