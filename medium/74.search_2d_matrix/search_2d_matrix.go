package main

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	for l < r {
		mid := (l + r) / 2
		if matrix[mid][len(matrix[0])-1] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	row := l
	l, r = 0, len(matrix[0])-1
	for l <= r {
		mid := (l + r) / 2
		if matrix[row][mid] == target {
			return true
		} else if matrix[row][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
