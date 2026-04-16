package main

func isStrictlyPalindromic(n int) bool {
	l, r := 2, n-2
	for l <= r {
		if !isPalindrome1(toBase(n, l)) {
			return false
		}
		l++
	}
	return true
}

func isPalindrome1(s []byte) bool {
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

func toBase(n, base int) []byte {
	if n == 0 {
		return []byte{0}
	}

	var digits []byte
	for n > 0 {
		digits = append(digits, byte(n%base))
		n /= base
	}

	// reverse (digits come out least-significant first)
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}
