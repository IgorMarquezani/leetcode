package main

import (
	"fmt"
	"slices"
)

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    slices.Sort(nums)
    result := 1
    count := 1

    for i := 0; i < len(nums)-1; i++ {
        for nums[i] == nums[i+1] && i < len(nums)-2 {
            i++
        }
        
        if nums[i] == nums[i+1]-1 {
            count++
            if count > result {
                result = count
            }
        } else {
            count = 1
        }
    }

    return result
}

func main() {
    v := []int{100,4,200,1,3,2}
    println(longestConsecutive(v))
    v = []int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 2}
    println(longestConsecutive(v))
    v = []int{}
    println(longestConsecutive(v))
    v = []int{0}
    println(longestConsecutive(v))
    v = []int{0, 0}
    println(longestConsecutive(v))
    v = []int{9,1,4,7,3,-1,0,5,8,-1,6}
    println(longestConsecutive(v))
}
