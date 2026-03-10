package main

func countStudents(students []int, sandwiches []int) int {
	studentPrefs := make(map[int]int)
	for _, s := range students {
		studentPrefs[s]++
	}

	for len(sandwiches) > 0 {
		val := studentPrefs[sandwiches[0]]
		if val > 0 {
			studentPrefs[sandwiches[0]]--
			sandwiches = sandwiches[1:]
		} else {
			return len(sandwiches)
		}
	}

	return 0
}
