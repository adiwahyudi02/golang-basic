package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("Adi")
	fmt.Println(result)
	
	fmt.Println(getHello("Wahyudi"))
}