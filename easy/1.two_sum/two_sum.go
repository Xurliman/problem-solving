package two_sum

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 && nums[0]+nums[1] == target {
		return []int{0, 1}
	}

	for firstKey, input := range nums {
		for secondKey, val := range nums[firstKey+1:] {
			if input+val == target {
				return []int{firstKey, secondKey + firstKey + 1}
			}
		}
	}
	return []int{0, 1}
}
