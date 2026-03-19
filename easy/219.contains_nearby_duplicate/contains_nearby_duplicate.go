package contains_nearby_duplicate

func containsNearbyDuplicate(nums []int, k int) bool {
	var hash = make(map[int]int)
	for i, n := range nums {
		pos, ok := hash[n]
		if !ok {
			hash[n] = i
			continue
		}

		if i-pos <= k {
			return true
		}

		hash[n] = i
	}

	return false
}
