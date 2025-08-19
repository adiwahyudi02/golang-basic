package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist) {
	if (blackList(name)) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blockFilter := func(name string) bool {
		return name == "Wahyudi"
	}

	registerUser("Adi", blockFilter)

	registerUser("Wahyudi", func(name string) bool {
		return name == "Wahyudi"
	})
}