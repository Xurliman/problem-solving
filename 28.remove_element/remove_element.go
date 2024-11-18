package _8_remove_element

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var spareKey = len(nums) - 1
	for key, num := range nums {
		if num == val {
			spareKey = getSpareKey(nums, val)
			if spareKey < key {
				continue
			}
			nums[key], nums[spareKey] = nums[spareKey], nums[key]
		}
	}
	return spareKey + 1
}

func getSpareKey(nums []int, val int) int {
	spareKey := len(nums)-1
	if spareKey >= 0 && nums[spareKey] == val {
		return getSpareKey(nums[:spareKey], val)
	}
	return spareKey
}