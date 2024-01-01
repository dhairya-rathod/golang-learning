package main

import "fmt"

func main() {
	/*
		-> In golang main function is an entrypoint. It takes no params also return nothing
		-> Go allow function to return multiple values
		-> Basic syntax
			func functionName(params) returnType {
				logic...
				return
			}
	*/

	greetUser("John")
	fmt.Println("My full name is ", fullName("John", "Doe"))

	flName, length := fullNameWithLength("Emily", "Backwood")
	fmt.Println(flName, length)

	variadicParam()
}

// returns nothing
func greetUser(name string) {
	fmt.Println("Hello, ", name)
}

// return string value
func fullName(fName, lName string) string {
	return fName + " " + lName
}

// multiple return
func fullNameWithLength(fName, lName string) (string, int) {
	flName := fullName(fName, lName)

	return flName, len(flName)
}
