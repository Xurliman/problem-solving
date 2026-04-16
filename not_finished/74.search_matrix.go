package not_finished

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1

	for l <= r {
		if matrix[(l+r)/2][0] < target {
			l++
		} else {
			r--
		}
	}

	fmt.Println(l, r)

	return true
}
