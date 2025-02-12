package roman_to_int

func romanToInt(s string) int {
	romans := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var (
		result int
		prev   int
	)

	for _, ch := range s {
		curr := romans[ch]
		if prev < curr {
			result -= prev
		} else {
			result += prev
		}
		prev = curr
	}
	return result + prev
}
