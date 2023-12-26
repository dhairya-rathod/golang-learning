package main

import "fmt"

// Like variable constant can be package or function level
const packageConst = 10

func main() {
	/*
		-> Values that never change during execution of program or are immutable
		-> Declared with the const keyword
		-> In Go constants can only hold values that the compiler can infer to be constant at the compile time
		-> You cannot make a runtime value constant
		-> Constants cannot be declared using the := syntax
	*/

	var number int = packageConst
	var decimalNumber float64 = packageConst
	// var stringNumber string = packageConst // throws an error

	// wrong example of constant
	/*
		var a = 10
		var b = 20
		const sum = a + b //a + b (value of type int) is not constant
	*/
	/*
		var a = 10
		var b = 20
		const sum = sum(a,b) //sum(a, b) (value of type int) is not constant
	*/

	fmt.Println("Constant values :: ", number, decimalNumber)
}

func sum(a, b int) int {
	return a + b
}
