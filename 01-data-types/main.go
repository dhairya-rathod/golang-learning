package main

import "fmt"

func main() {

	// ############ Data type String ############
	/*
		-> Double quotes and single quotes work differently in Go. Double quotes are used to denote strings, while single quotes are used to denote runes. A rune is a single character, while a string is a sequence of characters.
		-> var str string = "Hello"
		-> var r rune = 'H'
		-> var str string = "This is a \"quoted\" string." => "This is a "quoted" string."
	*/

	// Declaration - 01
	var myFirstName string //default zero value will be ""
	fmt.Println("My first name is ", myFirstName)
	myFirstName = "John" //assign a value
	fmt.Println("My first name is ", myFirstName)

	// Declaration - 02
	var myLastName string = "Doe" // define type on variable and assign value
	fmt.Println("My last name is ", myLastName)

	// Declaration - 03
	myFullName := "PhoenixDR" // Dynamic type declaration
	fmt.Println("My full name is ", myFullName)

	// multiple variable with same type declaration
	var a, b string
	var c, d string = "c", "d"
	fmt.Println("Random variables :: ", a, b, c, d)

	// ############ Data type Integer ############
	/*
		-> Declaration way will be same
		-> int store values from negative to positive
		-> uint(unsigned int) store values in positive only
		-> Basically we should use just int or uint and it'll take the type depending on system architecture
	*/

	// int8 and uint8 (2 raised to the power of 7)
	var int8 int8 = -128  // from -128 to 127
	var uint8 uint8 = 255 // from 0 to 255
	fmt.Println("Integer8 :: ", int8, uint8)

	// int16 and uint16 (2 raised to the power of 15)
	var int16 int16 = -32768  // from -32768 to 32767
	var uint16 uint16 = 65535 // from 0 to 65535
	fmt.Println("Integer16 :: ", int16, uint16)

	// int32 and uint32 (2 raised to the power of 31)
	var int32 int32 = -2147483648  // from -2147483648 to 2147483647
	var uint32 uint32 = 4294967295 // from 0 to 4294967295
	fmt.Println("Integer32 :: ", int32, uint32)

	// int64 and uint64 (2 raised to the power of 31)
	var int64 int64 = -9223372036854775808   // from -9223372036854775808 to 9223372036854775807
	var uint64 uint64 = 18446744073709551615 // from 0 to 18446744073709551615
	fmt.Println("Integer62 :: ", int64, uint64)

	// ############ Data type Float ############

	// float32
	var float32 float32 = 99.999888777666555444333222111
	fmt.Println("Float32 :: ", float32) // It'll print 99.999886 because of float32 and round of next decimals

	// float64
	var float64 float64 = 99.999888777666555444333222111
	fmt.Println("Float64 :: ", float64) // It'll print 99.99988877766656 because of float64 and round of next decimals

	revision()
}
