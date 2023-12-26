package main

import "fmt"

func main() {
	/*
		-> Slices are wrapper over array
		-> Slices have len and cap methods
		-> len to get the length of slice
		-> cap to get the capacity of slice
		-> GoLang has a concept of NIL which means slice is declared but not initialized
	*/

	// Slice declaration 01
	var slice1 []int // we do not specify length like array here
	fmt.Println("Check if slice is nil :: ", slice1 == nil)

	slice1 = append(slice1, 1) // append function to add elements in slice
	slice1 = append(slice1, 2, 3, 4, 5, 6)
	fmt.Println("Slice 1 :: ", slice1)

	// slice declaration 02
	var slice2 []int = make([]int, 5) //first arg is type, 2nd is length and third is capacity
	fmt.Println("Slice  :: ", slice2)

	// slice declaration 03
	slice3 := []int{1, 2, 3}
	fmt.Println("Slice 3 :: ", slice3)

	// length and capacity
	fmt.Println("length of slice3 :: ", len(slice3))
	fmt.Println("capacity of slice3 :: ", cap(slice3)) // capacity usually 25% or 50% of length
}
