package main

import "fmt"

/*
	Error Interface

	In Go, there is an interface called `error` which acts as a contract for error handling.
	Any type that implements the `Error() string` method automatically satisfies the `error` interface.
	This allows functions to return errors in a standardized way.
*/

/*
	Creating Errors

	In Go, we don’t need to create errors manually.
	The Go standard library already provides helpers for creating errors easily,
	which are available in the "errors" package.
*/

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}