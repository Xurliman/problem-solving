package main

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	start, end := -1, -1

	for left <= right {
		if nums[left] == target && start == -1 {
			start = left
		} else if nums[left] < target {
			left++
		}

		if nums[right] == target && end == -1 {
			end = right
		} else if nums[right] > target {
			right--
		}

		if start != -1 && end != -1 {
			break
		}
	}

	return []int{start, end}
}
