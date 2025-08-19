package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Counter = ", counter)
		counter++
	}

	fmt.Println("----------------")

	// the for has statement
	// the init statement
	// the condition statement
	// the post statement

	for i := 1; i <= 10; i++ {
		fmt.Println("Counter = ", i)
	}


	fmt.Println("----------------")

	// for range, can be used for interating over a data collection
	// like array, slice, map

	names := []string{"Adi", "Wahyudi"}
	for i, name := range names {
		fmt.Println(i, name)
	}
}