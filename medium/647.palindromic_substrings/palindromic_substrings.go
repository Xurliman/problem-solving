package main

import "fmt"

// abc: a, ab, abc, b, bc, c
func countSubstrings(s string) int {
	var count int
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			fmt.Println(s[i : j+1])
			if isPalindrome(s[i : j+1]) {
				count++
			}
		}
	}
	return count
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
