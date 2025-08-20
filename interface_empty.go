package main

import "fmt"

/*
	Empty Interface or Any

	Go is not an object-oriented programming language.
	In typical OOP languages, there is often a top-level parent type
	that can represent all data types in the language.
	For example, in Java there is java.lang.Object.

	To handle similar cases in Go, we can use the empty interface.
	An empty interface is an interface with no method declarations.
	This means that every data type automatically implements it.
	The empty interface also has a type alias called `any`.
*/

func Ups () any {
	// return 1
	// return true
	return "Ups"
}

func main() {
	var empty any = Ups()
	fmt.Println(empty)
}