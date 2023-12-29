package main

import "fmt"

func main() {
	ans01 := challenge01(5, 2)
	ans02 := challenge02(33)

	fmt.Println(ans01, ans02)
}

func challenge01(a, b int) int {
	// Write a program to multiply two number without using multiply sign
	ans := 0
	for i := 0; i < b; i++ {
		ans = ans + a
	}

	return ans
}

func challenge02(n int) []int {
	// Write the code which outputs prime numbers in the interval from 2 to n
	ans := []int{}
	index := 2

	for index <= n {
		index02 := 2
		isPrime := true

		for index02 < index {
			if index%index02 == 0 {
				isPrime = false
			}
			index02++
		}

		if isPrime {
			ans = append(ans, index)
		}
		index++
	}
	return ans
}
