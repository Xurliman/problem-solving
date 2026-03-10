package _2_longest_parentheses

func longestValidParentheses(s string) int {
	var left, right, maxR int
	for _, ch := range s {
		if string(ch) == "(" {
			left++
		} else {
			right++
		}

		if left == right {
			tmp := left * 2
			if maxR < tmp {
				maxR = tmp
			}
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "(" {
			left++
		} else {
			right++
		}

		if left == right {
			tmp := left * 2
			if maxR < tmp {
				maxR = tmp
			}
		} else if right < left {
			left, right = 0, 0
		}
	}

	return maxR
}
