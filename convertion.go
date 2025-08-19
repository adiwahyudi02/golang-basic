package main

import "fmt"

func main() {
	var value32 int32 = 32768
	var value64 int64 = int64(value32)

	// it will return -32768, because it's exceeds the range of int16
	// it's called number overflow, so it will return to the minimum value
	var value16 int16 = int16(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value16)


	var name = "Adi"
	var aByte = name[0]

	var aString = string(aByte)
	fmt.Println(name)
	fmt.Println(aByte)
	fmt.Println(aString)
}