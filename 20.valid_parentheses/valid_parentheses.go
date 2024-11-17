package valid_parentheses

func isValid(s string) bool {
	var stack []rune
	brackets := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		if openingBracket, ok := brackets[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != openingBracket {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
