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
		-> Functions in golang are First-class citizens (same as javascript) which means we can use it as a variable, can be return from another function and can be passed as an argument to another function
	*/

	greetUser("John")
	fmt.Println("My full name is ", fullName("John", "Doe"))

	flName, length := fullNameWithLength("Emily", "Backwood")
	fmt.Println(flName, length)

	// variadicParam()

	// first-class function
	fnEx := func() {
		println("Yo boiii")
	}
	fnEx()

	// function as argument
	sayNamaste("Emily", "Backwood", fullName)

	// currying function
	abc := curryingFn(7)(7)
	sumFive := curryingFn(5)
	ans := sumFive(4)
	ans2 := sumFive(10)
	println(abc, ans, ans2)
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

// function as a param
func sayNamaste(fName, lName string, fn func(string, string) string) {
	fullName := fn(fName, lName)

	fmt.Println("Namaste ", fullName)
}

// return function from a function
type fnType func(int) int // function type for code readability

func curryingFn(a int) fnType {
	return func(b int) int {
		return a + b
	}
}
