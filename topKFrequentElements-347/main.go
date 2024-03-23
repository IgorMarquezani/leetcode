package main

import (
	"fmt"
	"slices"
	"sort"
)

type Frequency struct {
	num, count int
}

type Frequencys []Frequency

func (f Frequencys) Len() int { return len(f) }

func (f Frequencys) Swap(i, j int) {
	f[j], f[i] = f[i], f[j]
}

func (f Frequencys) Less(i, j int) bool {
	return f[i].count > f[j].count
}

func topKFrequent(nums []int, k int) []int {
	var (
		result   = []int{}
		trakcing = Frequencys{}
	)

	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {
		trakcing = append(trakcing, Frequency{
			count: 1,
			num:   nums[i],
		})

		if i == len(nums)-1 {
			break
		}

		for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
			v := trakcing[trakcing.Len()-1]
			v.count++
			trakcing[trakcing.Len()-1] = v
		}
	}

	sort.Sort(trakcing)
	fmt.Println(trakcing)

	for i := 0; i < k; i++ {
		result = append(result, trakcing[i].num)
	}

	return result
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 9, 9, 9, 2, 3, 6, 4, 4}, 3))
}
