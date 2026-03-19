package majority_element

func majorityElement(nums []int) int {
	var hash = make(map[int]int)
	l := len(nums) / 2
	for _, n := range nums {
		hash[n]++

		if hash[n] >= l {
			return n
		}
	}

	return 0
}
