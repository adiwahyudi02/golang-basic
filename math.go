package main

import "fmt"

func main() {

	// math operator
	// + - * / %

	var a = 1
	var b = 2

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("a + b = ", a + b)
	fmt.Println("a - b = ", a - b)
	fmt.Println("a * b = ", a * b)
	fmt.Println("a / b = ", a / b)
	fmt.Println("a % b = ", a % b)

	fmt.Println("----------------")

	// augmented assignment operator
	// += -= *= /= %=
	a += b
	fmt.Println("a += b = ", a)
	a -= b
	fmt.Println("a -= b = ", a)
	a *= b
	fmt.Println("a *= b = ", a)
	a /= b
	fmt.Println("a /= b = ", a)
	a %= b
	fmt.Println("a %= b = ", a)


	fmt.Println("----------------")

	// unary operator
	// ++ --
	a++
	fmt.Println("a++ = ", a)
	a--
	fmt.Println("a-- = ", a)
}