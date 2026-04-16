package main

import "sort"

// 4,7,9
// 2,5,8,8
func minPairSum(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var m int
	l, r := 0, len(nums)-1
	for l <= r {
		sum := nums[l] + nums[r]
		if m < sum {
			m = sum
		}
		l++
		r--
	}
	return m
}
