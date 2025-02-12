package smallest_equal

func smallestEqual(nums []int) int {
	var smallest = -1
	for i, num := range nums {
		remainder := i % 10
		if remainder == num {
			return i
		}
	}
	return smallest
}
