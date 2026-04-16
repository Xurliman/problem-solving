package main

func appendCharacters(s string, t string) int {
	var count int
	var l, r int

	for r < len(t) {
		if l >= len(s) {
			count++
			r++
		}
		if l < len(s) && s[l] == t[r] {
			r++
		}
		l++
	}
	return count
}
