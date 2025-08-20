package main

import "fmt"

/*
	Pointers in Methods

	Although methods are associated with structs, the struct data accessed
	inside a method is passed by value by default.

	It is highly recommended to use pointers in methods to avoid unnecessary
	memory duplication when calling the method.

	Using pointers ensures that the method works with the original struct
	data and is more memory-efficient.
*/

type Man struct {
	Name string
}

// This won't change the Name, because the man is passed by value
// func (man Man) Merried() {
// 	man.Name = "Mr. " + man.Name
// }

func (man *Man) Merried() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{Name: "Adi"}

	man.Merried()
	fmt.Println(man.Name)
}