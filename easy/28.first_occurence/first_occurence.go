package first_occurence

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	haystackLen := len(haystack)
	for k, h := range haystack {
		if h == rune(needle[0]) {
			if k + needleLen > haystackLen {
				return -1
			} else if haystack[k:k + needleLen] == needle {
				return k
			} else {
				continue
			}
		}
	}
	return -1
}
