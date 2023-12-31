package main

import "fmt"

func main() {
	/*
		-> Some of the control flow statement supported by golang are break, continue, label and goto
		-> If we have nested loop label statement is used to manage outer loop
		-> goto statement allow us to jum to specific code block
	*/

	// mimic do while loop using break statement
	a := 10
	for {
		fmt.Println("Value of a :: ", a)
		a--
		if a == 5 {
			break
		}
	}

	// continue statement
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v is a odd number\n", i)
	}

	// label statement
outerForLoop:
	for i := 0; i < 3; i++ {
		// innerForLoop:
		for j := 0; j < 3; j++ {
			// for k := 0; k < 3; k++ {
			// if j == 2 && k == 2 {
			// 	continue innerForLoop
			// }
			if i == 1 && j == 1 /*&& k == 1*/ {
				continue outerForLoop
			}
			fmt.Printf("I is %v & j is %v\n", i, j)
			// }
		}
	}

	// goto statement
	abc := 5
	if abc == 5 {
		goto printA
	}

printA:
	fmt.Println("Value of abc is ::", abc)
}
