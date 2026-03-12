package main

import "strings"

func reverseVowels(s string) string {
	vowels := map[byte]struct{}{
		'a': struct{}{},
		'e': struct{}{},
		'i': struct{}{},
		'o': struct{}{},
		'u': struct{}{},
		'A': struct{}{},
		'E': struct{}{},
		'I': struct{}{},
		'O': struct{}{},
		'U': struct{}{},
	}

	var chars = make([]byte, 0, len(s))
	for i := range s {
		chars = append(chars, s[i])
	}

	var left, right = 0, len(s) - 1

	for left < right {
		if _, ok := vowels[chars[right]]; !ok {
			right--
			continue
		}

		if _, ok := vowels[chars[left]]; !ok {
			left++
			continue
		}

		chars[left], chars[right] = chars[right], chars[left]
		right--
		left++
	}

	var str strings.Builder
	str.Grow(len(s))
	for _, char := range chars {
		str.WriteByte(char)
	}
	return str.String()
}
