package surroundedRegions

import "fmt"

func SolveOld(board [][]byte) {
	m, n := len(board), len(board[0])
	copyBoard := make([][]byte, m)

	// 將所有元素先改成 X
	for i := 0; i < m; i++ {
		copyBoard[i] = make([]byte, n)
		copy(copyBoard[i], board[i])

		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}

	// 由邊界開始尋找沒有被包圍的島
	directs := []int{1, 0, -1, 0, 1}
	bfs := func(row, col int) {
		if copyBoard[row][col] == 'O' {
			queue := [][]int{{row, col}}

			for len(queue) > 0 {

				q := queue[0]
				queue = queue[1:]
				board[q[0]][q[1]] = 'O'
				copyBoard[q[0]][q[1]] = 'X'

				for s := 0; s < 4; s++ {
					r, c := q[0]+directs[s], q[1]+directs[s+1]

					if r > 0 && r < m && c > 0 && c < n && copyBoard[r][c] == 'O' {
						queue = append(queue, []int{r, c})
					}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		bfs(i, 0)
		bfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		bfs(0, j)
		bfs(m-1, j)
	}
}

func Solve(board [][]byte) {
	m, n := len(board), len(board[0])
	directs := []int{-1, 0, 1, 0, -1}

	bfs := func(row, col int) {
		if board[row][col] == 'O' {
			queue := [][]int{{row, col}}

			for len(queue) > 0 {
				r, c := queue[0][0], queue[0][1]
				queue = queue[1:]

				for s := 0; s < 4; s++ {
					nextR, nextC := r+directs[s], c+directs[s+1]
					if nextR > 0 && nextR < m && nextC > 0 && nextC < n && board[nextR][nextC] == 'O' {
						queue = append(queue, []int{nextR, nextC})
					}
				}

				board[r][c] = 'R'
			}
		}
	}

	for i := 0; i < m; i++ {
		bfs(i, 0)
		bfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		bfs(0, j)
		bfs(m-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'R' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func SolveDfs(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(int, int)
	dfs = func(row, col int) {
		if row >= 0 && row < m && col >= 0 && col < n && board[row][col] == 'O' {
			board[row][col] = 'R'
			dfs(row+1, col)
			dfs(row-1, col)
			dfs(row, col+1)
			dfs(row, col-1)
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	fmt.Printf("%s\n",board)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'R' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}
