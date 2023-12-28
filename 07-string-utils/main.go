package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello there")

	// Accessing byte data from string
	str := "Aello thera"
	firstChar := str[0]
	fmt.Println("Byte value of first char", firstChar)
	lastChar := str[len(str)-1]
	fmt.Println("Byte value of last char", lastChar)

	// convert byte or rune to string
	var char byte = 97
	var b2s string = string(char)
	fmt.Println("Byte to string", b2s)

	var char1 rune = 'A'
	var r2s string = string(char1)
	fmt.Println("Rune to string", r2s)

	// get substring(slicing a string, This is same as 06-slicing)
	greeting := "How are you?"
	firstWord := greeting[0:3]
	secondWord := greeting[4:8]
	fmt.Println("Words of greeting :: ", firstWord, secondWord)
}
