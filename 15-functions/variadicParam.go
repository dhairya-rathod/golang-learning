package main

import "fmt"

func variadicParam() {
	/*
		-> Go supports variadic params means it allow a function to accept any number of arguments
		-> it is similar to rest operator of javascript 🩷
		-> Basic syntax
			func functionName(restParams ...paramType) returnType {
				logic...
				return
			}
		-> here variadic params will be treated as slice
		-> When dealing with more than three params it is advisable to use struct
	*/

	fmt.Println("Sum of 1,2,3,4,5 is :: ", sum(1, 2, 3, 4, 5))

	printPersonDetails(personOption{
		fName: "Emily",
		lName: "Backwood",
		age:   21,
	})
}

func sum(numbers ...int) int {
	ans := 0
	for _, v := range numbers {
		ans = ans + v
	}
	return ans
}

// struct as a param
type personOption struct {
	fName string
	lName string
	age   int
}

func printPersonDetails(options personOption) {
	fmt.Printf("Full name is %v %v and age is %v\n", options.fName, options.lName, options.age)
}
