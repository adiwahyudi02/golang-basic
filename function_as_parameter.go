package main

import "fmt"

// we can also use type declaration for the function
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) string {
	filteredName := filter(name)

	return "Hello " + filteredName
}

func spamFilter (name string) string {
	if name == "Fool" {
		return "..."
	} else {
		return name
	}
}

func main() {
	helloPass := sayHelloWithFilter("Adi", spamFilter)
	fmt.Println(helloPass)

	helloSpam := sayHelloWithFilter("Fool", spamFilter)
	fmt.Println(helloSpam)
}