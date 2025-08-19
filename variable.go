package main

import "fmt"

func main() {
	// in var we have to always use the variable after we declare it
	// we need to define the type when we don't spesify the initial value
	var name string

	name = "Adi"
	fmt.Println(name)

	name = "Wahyudi"
	fmt.Println(name)

	// spesify the initial value, we don't need to define the type.
	// the type will be inferred from the initial value
	var country = "Indonesia"
	fmt.Println(country)

	// mostly used, we can use short declaration syntax
	hoby := "Coding"
	fmt.Println(hoby)

	hoby = "Sleeping"
	fmt.Println(hoby)


	// we can declare multiple variable
	var (
		firstName = "Adi"
		lastName = "Wahyudi"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}