package main

import "fmt"

func main() {
	/*
		-> Go has only for loop which is capable of replicating for, while and do-while loop
	*/

	// normal for loop like other programming language
	for i := 0; i < 1; i++ {
		fmt.Println("Index is  :: ", i)
	}

	divider()
	// for loop as while
	i := 0
	for i < 3 {
		fmt.Println("Index is  :: ", i)
		i++
	}

	divider()
	// for range loop
	arr := [5]int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Printf("Index is %v and Value is %v\n", index, value)
	}

	divider()
	str := "abcdefg"
	for index, value := range str {
		// here value will ve byte value of string char (99 to 103)
		fmt.Printf("Index is %v and Value is %v\n", index, value)
	}

	divider()
	myMap := map[string]string{
		"name": "name",
		"age":  "age",
		"addr": "addr",
	}
	for key, value := range myMap {
		fmt.Printf("Key is %v and Value is %v\n", key, value)
	}
}

func divider() {
	fmt.Printf("\n---------------\n")
}
