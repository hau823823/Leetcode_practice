package wordSearch

func Exist(board [][]byte, word string) bool {

	var backtrack func(int, int, int) bool 
	backtrack = func(row, col, idx int) bool {
		if idx == len(word) {
			return true
		}
		if row < 0 || row > len(board) - 1 || col < 0 || col > len(board[0]) - 1 || board[row][col] != word[idx] {
			return false
		}

		orgVal := board[row][col]
		board[row][col] = '*'
		res := backtrack(row, col-1, idx+1) || backtrack(row, col+1, idx+1) || backtrack(row-1, col, idx+1) || backtrack(row+1, col,idx+1)
		board[row][col] = orgVal

		return res
 	}

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			if backtrack(row, col, 0) {
				return true
			}
		}
	}
	return false
}