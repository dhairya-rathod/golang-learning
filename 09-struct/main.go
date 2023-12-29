package main

import "fmt"

func main() {
	/*
		-> In Go struct allow us define our own custom data type
		-> A struct is created using 'type' and 'struct' keyword (is is similar to 'type' in typescript)
		-> Basically struct is like objects in typescript
	*/

	// The above code defines a struct type called studentType with fields for name, age, and marks.
	// @property {string} name - The name property is a string that represents the name of the student.
	// @property {int} age - The "age" property is an integer that represents the age of a student.
	// @property {[]int} marks - The "marks" property is a slice of integers.
	type studentType struct {
		name  string
		age   int
		marks []int
	}

	// Creating a variable named `student` of struct/type `studentType` and initializing its fields with the provided values.
	var student = studentType{
		name:  "john",
		age:   11,
		marks: []int{11, 21, 31, 41, 51},
	}

	student_2 := studentType{
		name:  "Peter",
		age:   11,
		marks: []int{11, 21, 31, 41, 51},
	}

	// implicit assign (must follow order of keys)
	student_3 := studentType{
		"Emily",                   // name
		21,                        // age
		[]int{11, 21, 31, 41, 51}, // marks
	}

	fmt.Printf("%+v, %+v, %+v \n", student, student_2, student_3)

	// using the dot notation we can  modify and use struct value
	fmt.Println("Age of a student :: ", student.age)
	student.age = 21
	fmt.Println("Age of a student :: ", student.age)

	// anonymous struct struct{}{}
	pet := struct {
		category string
		breed    string
		age      float64
	}{category: "Dog", breed: "Street Dog", age: 0.6}
	fmt.Printf("Anonymous struct :: %+v", pet)
}
