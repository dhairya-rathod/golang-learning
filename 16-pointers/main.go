package main

import "fmt"

func main() {
	/*
		-> The allocation of memory is divided into tow segments Stack and Heap
		-> Stack manages function call execution
		-> Heap responsible for dynamic storage of a variable
		-> Go runtime has built in garbage collector which clean heap memory automatically

		-> Pointer is a variable that stores the memory address of a value
		-> Passing large struct as arguments in functions may lead to performance concerns. Pointer offer a way to mitigate this potential degradation.
	*/

	// pointer
	name := "John"
	nameAdd := &name

	fmt.Println(nameAdd)

	// dereferencing
	fmt.Println(*nameAdd)

	// increment without pointer
	a := 1
	incrementWoP := func(x int) {
		x++
	}
	incrementWoP(a)
	fmt.Println("Value of a :: ", a) // output = 1, because we pass argument by value

	// increment with pointer
	b := 1
	incrementWP := func(x *int) {
		*x++
	}
	incrementWP(&b)
	fmt.Println("Value of a :: ", b) // output = 2, because we pass memory address(pointer) and mutate the value over there

	// pointer example with struct
	type myData struct {
		name string
		age  int
	}
	me := myData{
		name: "John Doo",
		age:  21,
	}

	updateName := func(data *myData, updateName string) {
		data.name = updateName // for struct go manage dereferencing for us
	}

	updateName(&me, "John Doe")

	fmt.Println("Updated my data :: ", me)
}
