package count_binary_substrings

func countBinarySubstrings(s string) int {
	var (
		count, prev, cur, i = 0, 0, 1, 1
	)

	for i < len(s) {
		if s[i-1] != s[i] {
			count += min(prev, cur)
			prev = cur
			cur = 1
		} else {
			cur++
		}

		i++
	}
	count += min(prev, cur)
	return count
}
