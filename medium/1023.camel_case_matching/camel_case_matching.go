package main

var lowercase = map[rune]struct{}{
	'a': {}, 'b': {}, 'c': {}, 'd': {}, 'e': {}, 'f': {}, 'g': {},
	'h': {}, 'i': {}, 'j': {}, 'k': {}, 'l': {}, 'm': {}, 'n': {},
	'o': {}, 'p': {}, 'q': {}, 'r': {}, 's': {}, 't': {}, 'u': {},
	'v': {}, 'w': {}, 'x': {}, 'y': {}, 'z': {},
}

func camelMatch(queries []string, pattern string) []bool {
	var res []bool
	for _, query := range queries {
		res = append(res, match(query, pattern))
	}
	return res
}

func match(query string, pattern string) bool {
	patterns := splitPattern(pattern)
	queries := splitPattern(query)

	if len(queries) > len(patterns) {
		return false
	}

	var l, r int
	for l < len(patterns) && r < len(queries) {
		if startsWith(queries[r], patterns[l]) {
			l++
		}
		r++
	}

	return l == len(patterns)
}

func startsWith(query, pattern string) bool {
	if len(pattern) > len(query) {
		return false
	}

	var curQuery, curPattern int
	for curQuery < len(query) && curPattern < len(pattern) {
		if pattern[curPattern] == query[curQuery] {
			curPattern++
		}
		curQuery++
	}
	return curPattern == len(pattern)
}

func splitPattern(pattern string) []string {
	var res []string
	var start int
	var flag bool
	for i, p := range pattern {
		if _, ok := lowercase[p]; ok {
			continue
		}
		if !flag {
			flag = true
			continue
		}
		res = append(res, pattern[start:i])
		start = i
	}
	res = append(res, pattern[start:])
	return res
}
