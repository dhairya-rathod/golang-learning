package main

// a is a package level variable and can be accessible anywhere in package
var a = "Package level variable"

func main() {
	// b is a function level variable and can be accessible within this function only
	var b = "Function level variable"

	{
		// c is a block level variable and can be accessible in this block{} only
		var c = "Block{} level variable"

		// a, b, & c can be accessible here
	}

	// a & b can be accessible here
}

// only a is accessible here
// var c = c * b // compiler error :: b & c is not defined
