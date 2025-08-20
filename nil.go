package main

/*
	Nil

	In many programming languages, an uninitialized object automatically has a value of null or nil.
	In Go, when we create a variable of a certain type, it automatically gets a default value.
	However, Go also has a special value called `nil`, which represents an empty or zero value.

	Nil can only be used with certain data types, such as:
	- interface
	- function
	- map
	- slice
	- pointer
	- channel
*/

func NewMap(name string) map[string]string {
	if (name == "") {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("Adi")
	if data != nil {
		println(data["name"])
	} else {
		println("Data is nil")
	}
}