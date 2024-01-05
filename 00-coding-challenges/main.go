package main

import (
	"fmt"
	"sort"
)

func main() {
	ans01 := challenge01(5, 2)
	ans02 := challenge02(33)
	ans03 := challenge03(4, []int{1, 2, 3, 4, 5})

	fmt.Println(ans01, ans02, ans03)
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
