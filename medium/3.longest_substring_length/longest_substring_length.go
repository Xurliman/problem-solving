package main

func lengthOfLongestSubstring(s string) int {
	var (
		result int
		tmp    []rune
	)

	tmp = make([]rune, 0)
	for _, ch := range s {
		has, idx := hasRune(tmp, ch)
		if has {
			tmp = tmp[idx+1:]
		}

		tmp = append(tmp, ch)

		if len(tmp) > result {
			result = len(tmp)
		}
	}

	return result
}

func hasRune(s []rune, r rune) (bool, int) {
	for idx, ch := range s {
		if ch == r {
			return true, idx
		}
	}

	return false, -1
}
