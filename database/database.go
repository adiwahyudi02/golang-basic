package database

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

var connection string

func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}