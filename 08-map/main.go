package main

import "fmt"

func main() {
	/*
		-> Map datastructure in go is use to store key-value pair
		-> Same as slice if map is declared but not initialized it'll be NIL
		-> Syntax map[keyDataType]valueDataType{
			key: value
		}
		-> Syntax to get value :: mapName[keyName]
		-> Adding a new key-value in nil map can cause an error
	*/

	myMap1 := map[string]int{
		"john":  11,
		"peter": 21,
	}

	// add value in map
	myMap1["emily"] = 27

	// overwrite map key-value
	myMap1["john"] = 51

	// read from a map
	emilyAge := myMap1["emily"]

	fmt.Println(myMap1, emilyAge)

	// comma ok idiom
	nameAge := map[string]int{
		"john":  0,
		"peter": 21,
		"emily": 100,
	}
	gradeJohn, ok := nameAge["john"] // get boolean in ok based on key is exist or not => true
	println(gradeJohn, ok)
	gradeUkn, ok := nameAge["unknown"] // get boolean in ok based on key is exist or not => false
	println(gradeUkn, ok)
}
