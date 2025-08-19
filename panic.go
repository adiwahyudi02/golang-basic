package main

import "fmt"

// Panic
// The panic function is used to stop the program immediately.
// Panic is usually called when something unexpected or fatal happens during execution.
// Once panic is triggered, the program will terminate,
// but any defer functions will still be executed before the program fully stops.

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")
	}

	fmt.Println("Processing")
}

func main() {
	runApp(true)
}