package main

func pivotArray(nums []int, pivot int) []int {
	l := make([]int, 0)
	g := make([]int, 0)
	e := make([]int, 0)

	for _, n := range nums {
		if n > pivot {
			g = append(g, n)
		} else if n < pivot {
			l = append(l, n)
		} else {
			e = append(e, n)
		}
	}

	return append(append(l, e...), g...)
}
