package main

func calculateScore(instructions []string, values []int) int64 {
	n := len(values)
	executed := make(map[int]struct{}, n)

	var (
		score int64
		i     int
		val   string
	)

	for i < n {
		if _, ok := executed[i]; ok {
			break
		}
		val = instructions[i]
		executed[i] = struct{}{}

		if val == "add" {
			score += int64(values[i])
			i++
		} else {
			i += values[i]
		}
	}

	return score
}
