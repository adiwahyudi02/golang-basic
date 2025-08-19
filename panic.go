package main

import (
	"fmt"
)

// Panic
// The panic function is used to stop the program immediately.
// Panic is usually called when something unexpected or fatal happens during execution.
// Once panic is triggered, the program will terminate,
// but any defer functions will still be executed before the program fully stops.

// Recover
// The recover function is used to catch data from a panic.
// By using recover, the panic process is stopped, allowing the program to continue running.

func endApp() {
	fmt.Println("End App")

	message := recover()
	if message != nil {
		fmt.Println("Error message: ", message)
	}
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Unexpected error occurred")
	}

	fmt.Println("Processing")
}

func main() {
	runApp(true)

	fmt.Println("Program still continue...")
}