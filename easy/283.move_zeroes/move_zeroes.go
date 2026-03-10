package main

func moveZeroes(nums []int) {
	l := 0

	for r := 0; r < len(nums); r += 1 {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}
