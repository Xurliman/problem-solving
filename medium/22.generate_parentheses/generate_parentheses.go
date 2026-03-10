package _2_generate_parentheses

import "strings"

func generateParenthesis(n int) []string {
	var stack, result []string

	var backtrack func(int, int) // declare the type

	backtrack = func(openN, closedN int) {
		if openN == n && closedN == n {
			result = append(result, strings.Join(stack, ""))
			return
		}

		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closedN)
			stack = stack[:len(stack)-1]
		}

		if closedN < openN {
			stack = append(stack, ")")
			backtrack(openN, closedN+1)
			stack = stack[:len(stack)-1]
		}

	}

	backtrack(0, 0)
	return result
}
