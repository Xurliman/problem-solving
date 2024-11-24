package search_insert_position

func searchInsert(nums []int, target int) int {
	var pos int
	for k, num := range nums {
		if num == target {
			return k
		}
		if num > target {
			break
		}
		pos++
	}
	return pos
}
