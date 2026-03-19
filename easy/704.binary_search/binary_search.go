package binary_search

// nums = [-1,0,3,5,9,12], target = 9
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		if nums[(l+r)/2] == target {
			return (l + r) / 2
		}

		if nums[(l+r)/2] > target {
			r = (l+r)/2 - 1
		} else {
			l = (l+r)/2 + 1
		}
	}

	return -1
}
