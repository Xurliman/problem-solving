package main

func search(nums []int, target int) int {
	numsMap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}

	idx, ok := numsMap[target]
	if ok {
		return idx
	}
	return -1
}
