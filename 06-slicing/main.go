package main

import "fmt"

func main() {
	fmt.Println("Slicing")
	/*
		-> Go allow us to slice a slice or an array
		-> For e.x a[1:5]
		-> Here we'll get values from a's index 0 to 4 (NOT THE VALUE OF INDEX 5)
		-> Be careful with slicing as new slice/array after slicing share the shame reference
		which means manipulating anyone will get reflected on another
		-> To prevent this we can use built in copy function
	*/

	// Different slicing example

	// slicing from and to [From:To]
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[1:9]
	fmt.Println("Sliced a :: ", b, a) // [2 3 4 5 6 7 8 9]

	// slicing to [<blank>:To]
	c := a[:5]
	fmt.Println("Sliced a :: ", c) // [1, 2, 3, 4, 5]

	// slicing from [from:<blank>]
	d := a[5:]
	fmt.Println("Sliced a :: ", d) // [6 7 8 9 10]

	// first index to last index [:]
	e := a[:]
	fmt.Println("Sliced a :: ", e) // [1 2 3 4 5 6 7 8 9 10]

	// manipulating slices
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[:3]

	slice2[1] = 22
	fmt.Println("updated slices :: ", slice1, slice2) //[1 22 3 4 5] [1 22 3]

	slice2 = append(slice2, 33)
	fmt.Println("updated slices :: ", slice1, slice2) //[1 22 3 33 5] [1 22 3 33]

	// copy slice to new slice
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := make([]int, len(sliceA))
	copy(sliceB, sliceA[:3])

	fmt.Println("updated slices :: ", sliceA, sliceB) //[1 22 3 33 5] [1 22 3 33]
}
