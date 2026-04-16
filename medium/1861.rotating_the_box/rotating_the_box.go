package main

// | # | . | . | * | . | # | . |
func rotateTheBox(boxGrid [][]byte) [][]byte {
	n, m := len(boxGrid), len(boxGrid[0])
	for r := 0; r < n; r++ {
		empty := m - 1
		for c := m - 1; c >= 0; c-- {
			switch boxGrid[r][c] {
			case '*':
				empty = c - 1
			case '#':
				boxGrid[r][c] = '.'
				boxGrid[r][empty] = '#'
				empty--
			}
		}
	}

	res := make([][]byte, m)

	for c := 0; c < m; c++ {
		inner := make([]byte, n)
		for r := 0; r < n; r++ {
			inner[n-r-1] = boxGrid[r][c]
		}

		res[c] = inner
	}
	return res
}
