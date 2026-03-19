package word_pattern

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	hash := make(map[rune]string)
	used := make(map[string]struct{})

	for i, ch := range pattern {
		val, ok := hash[ch]
		if !ok {
			if _, isUsed := used[words[i]]; isUsed {
				return false
			}

			hash[ch] = words[i]
			used[words[i]] = struct{}{}
			continue
		}

		if val != words[i] {
			return false
		}
	}

	return true
}
