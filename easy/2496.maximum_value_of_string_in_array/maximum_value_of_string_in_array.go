package main

import "strconv"

func maximumValue(strs []string) int {
	var maxLen int
	for _, str := range strs {
		val, err := strconv.Atoi(str)
		if err != nil {
			val = len(str)
		}

		if val > maxLen {
			maxLen = val
		}
	}
	return maxLen
}
