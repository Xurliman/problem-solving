package length_of_last_word

func lengthOfLastWord(s string) int {
	sRunes := []rune(s)
	var length int
	for i := len(sRunes) - 1; i >= 0; i-- {
		if sRunes[i] == ' ' {
			if length == 0 {
				continue
			} else {
				break
			}
		}
		length++
	}
	return length
}
