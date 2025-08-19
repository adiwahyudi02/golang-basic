package main

import "fmt"

func main() {

	// Closure
	// The ability of a function to interact with variables 
	// from its surrounding scope, even after that scope has finished executing.
	// This makes closures very powerful for maintaining state, encapsulating logic, 
	// and creating more flexible code structures.
	// Please use closures wisely when building applications, as overusing them 
	// can make the code harder to read and debug.

	counter := 0

	increment := func() {
		counter++
	}

	increment()

	fmt.Println("Counter = ", counter)
}