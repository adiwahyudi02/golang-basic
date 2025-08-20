package helper

import "fmt"

/*
	Package

	A package is a way to organize the code in a Go program.
	By using packages, we can structure and tidy up our code.
	Essentially, a package corresponds to a folder (directory) in the operating system.
*/

/*
	Access Modifier

	In other programming languages, there are keywords to define access modifiers
	for functions or variables.
	In Go, access modifiers are determined simply by the case of the name:
	- If the name starts with an uppercase letter, it can be accessed from other packages.
	- If the name starts with a lowercase letter, it cannot be accessed from other packages.
*/

var version = "1.0.0"
var Application = "golang-basic"

func sayGoodBye(name string) string {
	return "Good bye " + name
}

func Example() {
	fmt.Println(Application)
	fmt.Println(version)
	fmt.Println(sayGoodBye("Adi"))
}

func SayHello(name string) string {
	return "Hello " + name
}