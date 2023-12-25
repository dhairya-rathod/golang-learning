package main

import "fmt"

func revision() {

	// Using package level variable
	packageVar = 12
	fmt.Println("Package level var value :: ", packageVar)

	// Declare a variable of type int and initialize it to the value 10.
	var myInt int = 10
	// Declare a variable of type string and initialize it to the value "Hello, world!".
	var myString string = "Hello, world!"
	// Declare a variable of type bool and initialize it to the value true.
	var myBool bool = true

	fmt.Println("Answers :: ", myInt, myString, myBool)
}
