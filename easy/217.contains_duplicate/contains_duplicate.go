package main

func containsDuplicate(nums []int) bool {
	mySet := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		_, ok := mySet[n]
		if ok {
			return true
		}

		mySet[n] = struct{}{}
	}

	return false
}
