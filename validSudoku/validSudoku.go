package validSudoku

// brute force
func IsValidSudoku(board [][]byte) bool {
	return CheckRow(board) && CheckCol(board) && Check3x3(board)
}

func CheckRow(board [][]byte) bool {

	for _, rows := range board {
		isDup := make([]bool, 9)
		for _, row := range rows {
			if row == '.' {
				continue
			} else if isDup[row-'1'] {
				return false
			} else {
				isDup[row-'1'] = true
			}
		}
	}
	return true
}

func CheckCol(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		isDup := make([]bool, 9)
		for j := 0; j < 9; j++ {
			col := board[j][i]
			if col == '.' {
				continue
			} else if isDup[col-'1'] {
				return false
			} else {
				isDup[col-'1'] = true
			}
		}
	}
	return true
}

func Check3x3(board [][]byte) bool {
	for i := 0; i < 9; i+=3 {
		for j :=0 ; j < 9; j+=3 {
			isDup := make([]bool, 9)
			for x := 0; x < 3; x++ {
				for y := 0; y < 3; y++ {
					val := board[i+x][j+y]
					if val == '.' {
						continue
					} else if isDup[val-'1'] {
						return false
					} else {
						isDup[val-'1'] = true
					}
				}
			}
		}
	}
	return true
}

// op
func IsValdiSudokuOp(board [][]byte) bool {
	var rows, cols, sqrs [9][9]bool

	for i, row := range board {
		for j, k := range row {
			if k == '.' {
				continue
			}
			if rows[i][k - '1'] || cols[j][k - '1'] || sqrs[i/3 * 3 + j/3][k - '1'] {
				return false
			}
			rows[i][k - '1'] = true
			cols[j][k - '1'] = true
			sqrs[i/3*3+j/3][k - '1'] = true
		}
	}

	return true
}