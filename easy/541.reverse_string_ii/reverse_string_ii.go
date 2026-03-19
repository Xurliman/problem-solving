package reverse_string_ii

import "strings"

func reverseStr(s string, k int) string {
	chunks := make([][]byte, len(s)/k+1)

	for i := 0; i < len(s); i++ {
		chunks[i/k] = append(chunks[i/k], s[i])
	}

	for pos := range chunks {
		if pos%2 == 0 {
			chunks[pos] = reverse(chunks[pos])
		}
	}
	var strs strings.Builder
	for _, ch := range chunks {
		for _, c := range ch {
			strs.WriteByte(c)
		}
	}
	return strs.String()
}

func reverse(s []byte) []byte {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
		continue
	}

	return s
}
