package main

import (
	"fmt"
)

func main() {
	/*
		-> Syntax ` if <CONDITION> {}
	*/
	isTen := 10
	if isTen == 10 {
		fmt.Println("Inside If block")
	} else if isTen == 20 {
		fmt.Println("Inside Else If block")
	} else {
		fmt.Println("Inside Else block")
	}

	// declaring a variable `age` with a value of 18 and then checking if the value of `age` is greater than 18.
	if age := 18; age > 18 {
		fmt.Println("User is adult")
	} else {
		fmt.Println("User is minor")
	}

	// fmt.Println(age) // age is not accessible here
}
