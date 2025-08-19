package main

import "fmt"

func main() {

	// Map Data Type
	// In an Array or Slice, we access data using numeric indexes starting from 0.
	// A Map is another data type that stores a collection of data,
	// but we can define the type of index (key) we want to use.
	//
	// Simply put, a Map is a collection of key-value pairs,
	// where the key must be unique and cannot be duplicated.
	//
	// Unlike Arrays and Slices, the number of data items we put into a Map
	// can be as many as we want, as long as the keys are different.
	// If we use the same key, the old value will automatically be replaced by the new one.



	// var person map[string]string = map[string]string{}
	// person["firstName"] = "Adi"
	// person["lastName"] = "Wahyudi"

	// or directly spesify the value
	person := map[string]string{
		"firstName": "Adi",
		"lastName": "Wahyudi",
	}

	fmt.Println(person)
	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])

	// if the key doesn't exist, it will return empty string (default value of string)
	fmt.Println(person["middleName"])


	fmt.Println("----------------")

	// Function Map
	// len(map) -> Get the number of elements in the map
	// map[key] -> Retrieve a value from the map by key
	// map[key] = value -> Update or insert a value in the map by key
	// make(map[TypeKey]TypeValue) -> Create a new map
	// delete(map, key) -> Remove an element from the map by key


	// book := map[string]string{
	// 	"title": "The Lord of the Rings",
	// 	"author": "J. R. R. Tolkien",
	// 	"year": "1954",
	// }

	// or we can like this
	book := make(map[string]string)
	book["title"] = "The Lord of the Rings"
	book["author"] = "J. R. R. Tolkien"
	book["year"] = "1954"
	book["wrong"] = "wrong"
	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)
}
