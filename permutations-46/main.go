package main

import (
	"fmt"
)

func factorial(n int) int {
	var result = 1

	for i := n; i > 0; i-- {
		result *= i

	}
	return result
}

func permute(nums []int) [][]int {
	result := make([][]int, 0)

	if len(nums) == 1 {
		cp := make([]int, 1)
		copy(cp, nums)
		return [][]int{{cp[0]}}
	}

	for i := 0; i < len(nums); i++ {
		v := make([]int, 0)
		cp := make([]int, len(nums))

		copy(cp, nums)
		v = append(v, cp[:i]...)
		v = append(v, cp[i+1:]...)

		perms := permute(v)

		for _, perm := range perms {
			perm = append(perm, nums[i])
			result = append(result, perm)
		}

		v = append(v, nums[i])
	}

	return result
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
