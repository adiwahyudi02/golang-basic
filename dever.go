package main

import "fmt"

// Defer in Go
// A defer statement is used to schedule a function call to be executed
// after the surrounding function finishes executing.
//
// Key characteristics of defer:
// 1. Deferred functions are always executed, even if the surrounding function
//    returns early or encounters a runtime error (panic).
// 2. Deferred calls are executed in LIFO (Last-In-First-Out) order.
// 3. It is often used for resource cleanup, such as closing files, unlocking mutexes,
//    or closing database connections.

func logging() {
	fmt.Println("Finished processing")
}

func runApplication() {
	defer logging()
	fmt.Println("Processing")
}

func main() {
	runApplication()
}