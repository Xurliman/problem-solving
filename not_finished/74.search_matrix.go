package not_finished

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1

	for l < r {
		if matrix[(l+r)/2][0] == target {
			return true
		}
	}

	return true
}
