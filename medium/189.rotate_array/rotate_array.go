package main

func rotate(nums []int, k int) []int {
	k %= len(nums)

	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}

	for i := 0; i < k/2; i++ {
		nums[i], nums[k-i-1] = nums[k-i-1], nums[i]
	}

	for i := k; i < (len(nums)+k)/2; i++ {
		nums[i], nums[len(nums)+k-i-1] = nums[len(nums)+k-i-1], nums[i]
	}

	return nums
}
