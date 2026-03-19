package sort_colors

func sortColors(nums []int) {
	hash := make(map[int]int)

	for _, num := range nums {
		hash[num]++
	}

	for i := 0; i < len(nums); i++ {
		if hash[0] > 0 {
			nums[i] = 0
			hash[0]--
		} else if hash[1] > 0 {
			nums[i] = 1
			hash[1]--
			continue
		} else {
			nums[i] = 2
			hash[2]--
		}
	}
}
