package main

import "fmt"

func main() {
	/*
		-> Each case operates with it's independent code block
		-> Like other languages break statement is not required at the end of each case in golang
		-> Here we can initialize variable with switch case
	*/

	month := "october"

	switch month {
	case "december":
		fmt.Println("It is december")
	case "october":
		fmt.Println("It is october")
	default:
		fmt.Println("Which month it is?")
	}

	//  initialize new variable
	switch day := "friday"; day {
	case "sunday":
		fmt.Println("Sunday it is 😁")
	case "friday":
		fmt.Println("Friday it is 🍺")
	default:
		fmt.Println("🤔")
	}

	// switch with multiple cases
	switch char := "a"; char {
	case "a", "b":
		println("char is a or b")
	case "c":
		println("char is c")
	default:
		println("char is c")
	}
}
