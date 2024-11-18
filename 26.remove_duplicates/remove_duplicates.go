package remove_duplicates

func removeDuplicates(nums []int) int {
	set := make(map[int]int)
	i := 0
	for _, num := range nums {
		if _, ok := set[num]; !ok {
			i++
		}
		set[num] = i
	}

	for k, s := range set {
		nums[s-1] = k
	}
	return len(set)
}
