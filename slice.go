package main

import "fmt"

func main() {

	// SLICE
	//
	// Slice Data Type
	// -----------------------------------------
	// - A Slice is a segment (portion) of an Array.
	// - Similar to an Array, but the size of a Slice can change (dynamic).
	// - A Slice and its underlying Array are always connected,
	//   meaning a Slice provides access to part or all of an Array.
	//
	// 
	// Slice Details
	// -----------------------------------------
	// A Slice consists of 3 main components:
	// 1. Pointer   -> points to the first element of the underlying Array.
	// 2. Length    -> the number of elements in the Slice.
	// 3. Capacity  -> the total number of elements from the Pointer
	//                 position to the end of the Array.
	//    * Note: length <= capacity
	//
	// 
	// Creating a Slice from an Array
	// -----------------------------------------
	// array[low:high] -> creates a Slice from index low up to (but not including) index high.
	// array[low:]     -> creates a Slice from index low to the end of the Array.
	// array[:high]    -> creates a Slice from index 0 up to (but not including) index high.
	// array[:]        -> creates a Slice from index 0 to the end of the Array.
	//
	// =========================================

	names := [...]string{"Dada", "Didi", "Dodo", "Dudu", "Dede", "Dadi"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	// it's same as slice4, 
	var slice5 []string = names[:]
	fmt.Println(slice5)

	fmt.Println("----------------")



	// SLICE FUNCTIONS
	// -----------------------------------------
	// len(slice)	-> Get the length of the slice
	// cap(slice) -> Get the capacity of the slice
	// append(slice, data) -> Create a new slice by adding data to the last position of the slice;
	//                        if the capacity is full, a new array will be created
	// make([]DataType, length, capacity) -> Create a new slice
	// copy(destination, source) -> Copy slice from source to destination


	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	daySlice1 := days[5:] // [Saturday Sunday]
	fmt.Println(daySlice1)

	daySlice1[0] = "New Saturday"
	daySlice1[1] = "New Sunday"
	fmt.Println(daySlice1) // [New Saturday New Sunday]
	// this mean, when we change the value of daySlice1, the value of days will also change
	fmt.Println(days) // [Monday Tuesday Wednesday Thursday Friday New Saturday New Sunday]

	fmt.Println("----------------")

	daySlice2 := append(daySlice1, "New Holiday")
	// because of the capacity is full, a new array will be created
	// this is the underlying array of daySlice2
	// daysBaru := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "New Saturday", "New Sunday", "New Holiday"}

	fmt.Println(daySlice1) // [New Saturday New Sunday]
	fmt.Println(daySlice2) // [New Saturday New Sunday New Holiday]

	// the days will not change, not added "New Holiday"
	fmt.Println(days) // [Monday Tuesday Wednesday Thursday Friday New Saturday New Sunday]

	fmt.Println("----------------")

	daySlice2[0] = "Old Saturday"
	// the daySlice1 will not change, because daySlice2 now has a new underlying array
	fmt.Println(daySlice1) // [New Saturday New Sunday]
	fmt.Println(daySlice2) // [Old Saturday New Sunday New Holiday]
	// the days will not change
	fmt.Println(days) // [Monday Tuesday Wednesday Thursday Friday New Saturday New Sunday]


	fmt.Println("----------------")

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Adi"
	newSlice[1] = "Wahyu"
	// newSlice[2] = "Di" // error, we should use append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) // 2
	fmt.Println(cap(newSlice)) // 5

	// newSlice2 still use the same underlying array, because the capacity is enough
	newSlice2 := append(newSlice, "Di")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2)) // 3
	fmt.Println(cap(newSlice2)) // 5


	// let's try to change the newSlice2
	newSlice2[0] = "Dadang"
	fmt.Println(newSlice) // [Dadang Wahyu]
	fmt.Println(newSlice2) // [Dadang Wahyu Di]


	fmt.Println("----------------")

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)



	fmt.Println("----------------")
	
	// Caution, the different of slice and array declaration

	thisArray := [...]int{1, 2, 3}
	thisSlice := []int{1, 2, 3}

	fmt.Println(thisArray) // [1 2 3]
	fmt.Println(thisSlice) // [1 2 3]
}