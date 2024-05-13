package main

import (
	"fmt"
	"sort"
)

func main() {
	ans01 := challenge01(5, 2)
	ans02 := challenge02(33)
	ans03 := challenge03(4, []int{1, 2, 3, 4, 5})
	ans04 := challenge04("anina")

	fmt.Println(ans01, ans02, ans03, ans04)
	fmt.Println(isPalindrome(121))
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

func challenge03(target int, input []int) []int {
	// two sum
	// sort the array
	sort.Ints(input)
	point1, point2 := 0, len(input)-1

	for point1 < point2 {
		if input[point1]+input[point2] == target {
			return []int{input[point1], input[point2]}
		}
		if input[point1]+input[point2] > target {
			point2--
		} else {
			point1++
		}

	}
	return []int{}
}

func challenge04(s string) bool {
	// check for palindrome string
	lChar, rChar := 0, len(s)-1

	shouldSkip := func(c byte) bool {
		return (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') || c == ' '
	}

	for lChar <= rChar {
		if shouldSkip(s[lChar]) {
			lChar++
			continue
		}
		if shouldSkip(s[rChar]) {
			rChar--
			continue
		}

		if s[lChar] != s[rChar] {
			return false
		}

		lChar++
		rChar--
	}

	return true
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	xStr := string(rune(x))
	lChar, rChar := 0, len(xStr)-1

	fmt.Println(xStr[lChar],xStr[rChar])
	for lChar <= rChar {
		if xStr[lChar] != xStr[rChar] {
			return false
		}

		lChar++
		rChar--
	}

	return true
}
