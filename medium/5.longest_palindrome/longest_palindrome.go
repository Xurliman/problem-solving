package main

func longestPalindrome(s string) string {
	var (
		res    string
		resLen int
		l, r   int
	)

	for i := 0; i < len(s); i++ {
		l, r = i, i
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if resLen < r+1-l {
				res = s[l : r+1]
				resLen = r + 1 - l
			}
			l -= 1
			r += 1
		}

		l, r = i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if resLen < r+1-l {
				res = s[l : r+1]
				resLen = r + 1 - l
			}
			l -= 1
			r += 1
		}
	}
	return res
}
