package main

import "fmt"

func main() {

	// Boolean operator
	// && is and
	// || is or
	// ! is not

	fmt.Println("True && True = ", true && true)
	fmt.Println("True && False = ", true && false)
	fmt.Println("False && True = ", false && true)
	fmt.Println("False && False = ", false && false)
	fmt.Println("True || True = ", true || true)
	fmt.Println("True || False = ", true || false)
	fmt.Println("False || True = ", false || true)
	fmt.Println("False || False = ", false || false)
	fmt.Println("!True = ", !true)
	fmt.Println("!False = ", !false)

	fmt.Println("----------------")

	var value = 90
	var attendance = 80

	var passValue = value > 80
	var passAttendance = attendance > 80

	var pass = passValue && passAttendance

	fmt.Println("passValue = ", passValue)
	fmt.Println("passAttendance = ", passAttendance)
	fmt.Println("pass = ", pass)
}