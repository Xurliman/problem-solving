package _758_min_operations

func minOperations(s string) int {
	var sim1, sim2 int

	for i, ch := range s {
		if (ch == '1' && i%2 == 0) || (ch == '0' && i%2 == 1) {
			sim1++
		}
		if (ch == '1' && i%2 == 1) || (ch == '0' && i%2 == 0) {
			sim2++
		}
	}

	return len(s) - max(sim1, sim2)
}
