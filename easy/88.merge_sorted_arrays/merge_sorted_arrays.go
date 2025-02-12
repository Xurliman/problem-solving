package merge_sorted_arrays

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}
	nums3 := make([]int, m)
	copy(nums3, nums1[:m])
	var i, j int
	for i < m && j < n {
		if nums3[i] < nums2[j] {
			nums1[i+j] = nums3[i]
			i++
		} else {
			nums1[i+j] = nums2[j]
			j++
		}
	}

	for i+j < m+n && len(nums2[j:]) >0 {
		nums1[i+j] = nums2[j]
		j++
	}
	for i+j < m+n && len(nums3[i:]) >0 {
		nums1[i+j] = nums3[i]
		i++
	}
}
