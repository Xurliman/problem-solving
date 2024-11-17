package palindrome_number

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	newStrRunes := make([]rune, length)
	for k, s := range str {
		newStrRunes[length-k-1] = s
	}

	newStr := string(newStrRunes)
	if newStr == str {
		return true
	}
	return false
}
