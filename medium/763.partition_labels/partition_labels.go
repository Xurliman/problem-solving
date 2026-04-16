package main

// 0 1 2 3 4 5 6 7 8 9 10
// a b a b c b a c a d e f e g d e h i j h k l i j
func partitionLabels(s string) []int {
	var hash = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := hash[s[i]]; ok {
			continue
		}

		hash[s[i]] = getRightMost(s, s[i])
	}

	var output = make([]int, 0)
	size, end := 0, 0
	for i := 0; i < len(s); i++ {
		size++
		if hash[s[i]] > end {
			end = hash[s[i]]
		} else if i == end {
			output = append(output, size)
			size, end, i = 0, 0, end
		}
	}
	return output
}

func getRightMost(s string, ch byte) int {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == ch {
			return i
		}
	}
	return 0
}
