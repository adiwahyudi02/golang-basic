package main

import (
	"fmt"
	"golang-basic/helper"
)

func main() {
	result := helper.SayHello("Adi")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // undefined: helper.version
	// fmt.Println(helper.sayGoodBye("Adi")) // undefined: helper.sayGoodBye
}