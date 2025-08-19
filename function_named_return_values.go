package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "Adi"
	lastName = "Wahyudi"

	return firstName, lastName
}

func main() {
	firstName, lastName := getCompleteName()
	fmt.Println(firstName, lastName)
}