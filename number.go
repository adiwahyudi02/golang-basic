package main

import "fmt"

// Number type in go

// INTEGER
// int8: -128 to 127
// int16: -32768 to 32767
// int32: -2147483648 to 2147483647
// int64: -9223372036854775808 to 9223372036854775807

// UNSIGNED INTEGERS
// uint8: 0 to 255
// uint16: 0 to 65535
// uint32: 0 to 4294967295
// uint64: 0 to 18446744073709551615

// FLOAT
// float32: -3.402823466e+38 to 3.402823466e+38
// float64: -1.7976931348623157e+308 to 1.7976931348623157e+308

// COMPLEX
// complex64: real: -3.402823466e+38 to 3.402823466e+38, imaginary: -3.402823466e+38 to 3.402823466e+38
// complex128: real: -1.7976931348623157e+308 to 1.7976931348623157e+308, imaginary: -1.7976931348623157e+308 to 1.7976931348623157e+308

// ALIASES
// byte: (uint8) 0 to 255
// rune: (int32) -2147483648 to 2147483647
// int: at least 32-bit signed integer (size depends on system: 32-bit or 64-bit)
// uint: at least 32-bit unsigned integer (size depends on system: 32-bit or 64-bit)

func main() {
	fmt.Println("One = ", 1)
	fmt.Println("Two = ", 2)
	fmt.Println("Three point five = ", 3.5)
}