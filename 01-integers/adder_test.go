package integers

import (
	// "fmt"
	"testing"
)

// func Add(x, y int) int {
// 	return x + y
// }

// func ExampleAdd() {
// 	sum := Add(1, 5)
// 	fmt.Println(sum)
// 	// Output: 6
// }

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4
	if sum != expected {
		t.Errorf("Expected '%d' but got '%d", expected, sum)
	}
}