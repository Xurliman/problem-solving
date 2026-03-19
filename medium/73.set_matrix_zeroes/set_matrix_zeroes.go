package set_matrix_zeroes

func setZeroes(matrix [][]int) {
	hashI := make(map[int]struct{})
	hashJ := make(map[int]struct{})

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				continue
			}

			hashI[i] = struct{}{}
			hashJ[j] = struct{}{}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			_, ok := hashI[i]
			if ok {
				matrix[i][j] = 0
			}

			_, ok = hashJ[j]
			if ok {
				matrix[i][j] = 0
			}
		}
	}
}
