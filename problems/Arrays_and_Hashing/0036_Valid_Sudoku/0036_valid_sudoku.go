package problems

func isValidSudoku(board [][]byte) bool {
	rows := make(map[int]map[byte]bool)
	cols := make(map[int]map[byte]bool)
	sqrs := make(map[[2]int]map[byte]bool)

	for r := 0; r < 9; r++ {
		rows[r] = make(map[byte]bool)

		for c := 0; c < 9; c++ {

			if cols[c] == nil {
				cols[c] = make(map[byte]bool)
			}

			if board[r][c] == '.' {
				continue
			}

			val := board[r][c]
			key := [2]int{r / 3, c / 3}

			if sqrs[key] == nil {
				sqrs[key] = make(map[byte]bool)
			}

			if rows[r][val] || cols[c][val] || sqrs[key][val] {
				return false
			}

			rows[r][val] = true
			cols[c][val] = true
			sqrs[key][val] = true
		}
	}

	return true
}
