package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, str := range strs {
		if len(prefix) > len(str) {
			prefix = prefix[:len(str)]
		} else {
			str = str[:len(prefix)]
		}
		prefix = getPrefix(prefix, str)
	}
	return prefix
}

func getPrefix(prefix string, word string) string {
	prefixLen := len(prefix)
	if prefix == word[0:prefixLen] {
		return prefix
	} else {
		return getPrefix(prefix[0:prefixLen-1], word[0:prefixLen-1])
	}
}
