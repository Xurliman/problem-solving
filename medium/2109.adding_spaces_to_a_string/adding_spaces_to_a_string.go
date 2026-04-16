package main

import "strings"

func addSpaces(s string, spaces []int) string {
	var res strings.Builder
	var start int
	for _, space := range spaces {
		res.WriteString(s[start:space])
		res.WriteString(" ")
		start = space
	}
	res.WriteString(s[start:])
	return res.String()
}
