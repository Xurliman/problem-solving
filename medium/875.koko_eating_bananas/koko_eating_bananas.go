package main

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, getMax(piles)
	mid := (left + right) / 2

	hours := countHours(piles, mid)
	for left < right {
		if hours > h {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (left + right) / 2
		hours = countHours(piles, mid)
	}
	return mid
}

func countHours(piles []int, k int) int {
	var i int
	var h int
	var cur = piles[i]
	for i < len(piles) {
		h += cur / k
		if cur%k > 0 {
			h++
		}
		i++
		if i == len(piles) {
			break
		}
		cur = piles[i]
	}
	return h
}

func getMax(piles []int) int {
	var m = piles[0]
	for i := 1; i < len(piles); i++ {
		if m < piles[i] {
			m = piles[i]
		}
	}
	return m
}
