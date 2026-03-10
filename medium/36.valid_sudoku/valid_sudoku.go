package main

func isValidSudoku(board [][]byte) bool {
	type pair struct {
		r, c int
	}

	l := len(board)
	rows := make([]map[byte]struct{}, l)
	cols := make([]map[byte]struct{}, l)
	squares := make(map[pair]map[byte]struct{}, l)

	for i := 0; i < l; i++ {
		rows[i] = make(map[byte]struct{}, l)
		cols[i] = make(map[byte]struct{}, l)
	}

	for r := 0; r < l; r++ {
		for c := 0; c < l; c++ {
			if board[r][c] == '.' {
				continue
			}

			if _, ok := rows[r][board[r][c]]; ok {
				return false
			}
			rows[r][board[r][c]] = struct{}{}

			if _, ok := cols[c][board[r][c]]; ok {
				return false
			}
			cols[c][board[r][c]] = struct{}{}

			if _, ok := squares[pair{r / 3, c / 3}]; !ok {
				squares[pair{r / 3, c / 3}] = make(map[byte]struct{}, l)
			}
			if _, ok := squares[pair{r / 3, c / 3}][board[r][c]]; ok {
				return false
			}

			squares[pair{r / 3, c / 3}][board[r][c]] = struct{}{}
		}
	}

	return true
}
