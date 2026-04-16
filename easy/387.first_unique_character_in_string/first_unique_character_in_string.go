package main

func firstUniqChar(s string) int {
	var hash = make(map[byte]struct{})
	var first int
	for i := range s {
		if _, ok := hash[s[i]]; ok && s[first] == s[i] {
			first++
		}
		hash[s[i]] = struct{}{}
	}

	return first
}
