package main

func minimumSteps(s string) int64 {
	var place int
	var res int

	for i := range s {
		if s[i] == '1' {
			continue
		}

		res += i - place
		place++
	}
	return int64(res)
}
