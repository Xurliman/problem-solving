package main

func intersection(nums1 []int, nums2 []int) []int {
	hash1 := make(map[int]bool)
	for _, n := range nums1 {
		hash1[n] = false
	}

	var res = make([]int, 0)
	for _, n := range nums2 {
		if visited, ok := hash1[n]; ok && !visited {
			res = append(res, n)
			hash1[n] = true
		}
	}

	return res
}
