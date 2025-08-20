package main

/*
	Package Initialization

	In Go, when we create a package, we can also define a special function
	that will be executed automatically when the package is initialized.

	This is useful in cases such as packages containing functions
	to communicate with a database — we can use an initialization
	function to open a connection as soon as the package is imported.

	To achieve this, we simply create a function named `init`.
	Go will automatically call `init()` when the package is imported.
*/

/*
	Blank Identifier (_)

	Sometimes we only want to run the `init` function in a package
	without actually using any of its exported functions, variables, or types.

	By default, Go will throw a compile-time error if a package is imported
	but not used anywhere in the code.

	To handle this, we can use the blank identifier `_` before the package name
	during import. This tells the compiler that we are importing the package
	only for its side effects (like running its `init` function).
*/

import (
	"fmt"
	"golang-basic/database"
	_ "golang-basic/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}