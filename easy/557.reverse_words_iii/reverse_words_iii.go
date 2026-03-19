package reverse_words_iii

import "strings"

func reverseWords(s string) string {
	words := strings.Split(s, " ")

	for i := 0; i < len(words); i++ {
		words[i] = string(reverse(splitter(words[i])))
	}

	return strings.Join(words, " ")
}

func splitter(s string) []rune {
	bytes := make([]rune, 0, len(s))
	for _, ch := range s {
		bytes = append(bytes, ch)
	}

	return bytes
}

func reverse(s []rune) []rune {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
		continue
	}

	return s
}
