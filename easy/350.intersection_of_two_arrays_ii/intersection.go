package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var left, right int
	var res = make([]int, 0)
	for left < len(nums1) && right < len(nums2) {
		if nums1[left] == nums2[right] {
			res = append(res, nums1[left])
			left++
			right++
			continue
		}

		if nums1[left] < nums2[right] {
			left++
		} else {
			right++
		}
	}

	return res
}
