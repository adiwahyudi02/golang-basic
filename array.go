package main

import "fmt"

func main() {
	// we have to spesify the size of array
	// after we spesicify the size of array, we can't change the size of array
	// in golang we can't remove the item of array, instead we can set the value to 0 for int or "" for string etc
	var names [2]string

	names[0] = "Adi"
	names[1] = "Wahyudi"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])

	// error, invalid argument: index 2 out of bounds [0:2]
	// names[2] = "Dadang"


	// we can declare array in one line
	// the type become optional, it's inferred from the initial value
	var values = [3]int{70, 80}

	// return [70 80 0]
	// by default, if the initial value is not spesified and the type of item is number, the value will be 0
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// we can undeclare the size of array by using "..."
	// but, only works when the initial value is defined directly
	var values2 = [...]int{
		70,
		80,
		90,
		100,
		110,
	}

	fmt.Println(values2)
	fmt.Println(values2[0])
	fmt.Println(values2[1])
	fmt.Println(values2[2])
	fmt.Println(values2[3])
	fmt.Println(values2[4])

	// function in array
	// len(array) => length of array
	// array[index] => value at index
	// array[index] = value => set value at index

	fmt.Println(len(values2))
	fmt.Println(values2[0])
	values2[0] = 100
	fmt.Println(values2[0])
}