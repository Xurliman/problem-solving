package valid_palindrome_ii

// abcdceba
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] == s[r] {
			l++
			r--
			continue
		}
		return false
	}

	return true
}
