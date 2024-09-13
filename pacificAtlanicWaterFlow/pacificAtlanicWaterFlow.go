package pacificAtlanicWaterFlow

func PacificAtlantic(heights [][]int) [][]int {
	res := [][]int{}
	row, col := len(heights), len(heights[0])

	pacific := make([][]bool, row)
	atlantic := make([][]bool, row)
	for r := 0; r < row; r++ {
		pacific[r] = make([]bool, col)
		atlantic[r] = make([]bool, col)
	}

	var dfs func(int, int, int, [][]bool)
	dfs = func(r, c, preHeight int, visit [][]bool) {
		if r < 0 || r >= row || c < 0 || c >= col || heights[r][c] < preHeight || visit[r][c] {
			return
		}
		visit[r][c] = true
		dfs(r+1, c, heights[r][c], visit)
		dfs(r-1, c, heights[r][c], visit)
		dfs(r, c+1, heights[r][c], visit)
		dfs(r, c-1, heights[r][c], visit)
	}

	for c := 0; c < col; c++ {
		dfs(0, c, heights[0][c], pacific)
		dfs(row-1, c, heights[row-1][c], atlantic)
	}

	for r := 0; r < row; r++ {
		dfs(r, 0, heights[r][0], pacific)
		dfs(r, col-1, heights[r][col-1], atlantic)
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}
