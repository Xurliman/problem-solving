package assign_cookies

import "sort"

func findContentChildren(g []int, s []int) int {
	var contentN int
	var left, right int

	sort.Ints(g)
	sort.Ints(s)

	for left < len(g) && right < len(s) {
		if g[left] <= s[right] {
			contentN++
			right++
			left++
			continue
		}

		for right < len(s) && g[left] > s[right] {
			right++
		}
	}

	return contentN
}
