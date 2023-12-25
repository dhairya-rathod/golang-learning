package main

import "fmt"

// package level variable. Can be accessible anywhere in the package
var packageVar int = 11

func main() {
	/*
		-> There can be package level variables or function level variables (In other words scope of the variable can be package or function)
		-> Variable declaration can be grouped with braces ()
		-> Multiple variables can be grouped in single line
		Examples:
		var age int
		var age int = 10
		var age = 10
		var age, name = 10, "John"
		age := 10
		age, name := 10, "John"
	*/

	fmt.Println("Package level var value :: ", packageVar)

	// Group similar variables
	var (
		age  int    = 21
		name string = "John"
	)
	fmt.Printf("%s is %d years old \n", name, age)

	// Explicit type assignment (i.e. assign type with variable declaration)
	var myAge int = 21
	var myName string = "John"
	fmt.Printf("%s is %d years old \n", myName, myAge)

	// Implicit type assignment (i.e. variable infer type from value)
	var yourAge = 51
	var yourName = "Peter"
	fmt.Printf("%s is %d years old \n", yourName, yourAge)

	// Shorthand variable declaration
	fullName := "John Doe"
	city := "NYC"
	year := 2023

	// city := "Taxas" //throw an error because can't update value with := unless there is new var on left side
	city, partnerName := "Taxas", "Jane Doe"
	fmt.Printf("%s was in %s since %d with %s \n", fullName, city, year, partnerName)

	revision()
}
