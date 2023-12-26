package main

import "fmt"

func main() {
	/*
		-> Syntax :: var <variable name> [length of an array]type of an array
	*/
	// Array declaration 01
	var arrDec1 [5]int
	fmt.Println("Array Dec 01 :: ", arrDec1) // all the values of this array will be zero

	// Array declaration 02
	var arrDec2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array Dec 02 :: ", arrDec2)

	// Array declaration 03
	arrDec3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array Dec 03 :: ", arrDec3)

	// Array declaration 04
	arrDec4 := [...]int{1, 2, 3, 4, 5} // here size of array will be dynamic
	fmt.Println("Array Dec 04 :: ", arrDec4)

	// Array Initialize 01
	var arrInit1 [5]int
	arrInit1[0] = 1
	arrInit1[4] = 5
	fmt.Println("Array Init 01 :: ", arrInit1) // value of 1,2,3 index will be zero

	// Array Initialize 02 (set value of index by <I>:<V>)
	var arrInit2 [5]int = [5]int{1, 2: 2, 5}   // => [1,0,2,5,0]
	fmt.Println("Array Init 02 :: ", arrInit2) // value of 1 & 5 index will be zero

	// get value of array index
	fmt.Println("Value at index 2 of arrDec2 :: ", arrDec2[2])

	// len function of array
	fmt.Println("Length of arrDec2 :: ", len(arrDec2))

	// Try to access out of range array element
	// fmt.Println("Value at index 5 of arrDec2 :: ", arrDec2[5]) // this will give error `invalid argument: index 5 out of bounds [0:5]`

	// Two dimensional array
	twoDArr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("2D array :: ", twoDArr)

	// Manipulate array elements
	arrDec2[2] = 12
	twoDArr[0][2] = 12
	fmt.Println("Array Dec 02 :: ", arrDec2)
	fmt.Println("2D array :: ", twoDArr)
}
