package rottingOranges

// bfs
func OrangesRottingBFS(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	directs := []int{-1, 0, 1, 0, -1}

	queue := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	step := 0
	for len(queue) > 0 {
		for _, orange := range queue {
			queue = queue[1:]
			for k := 0; k < 4; k++ {
				row, col := orange[0]+directs[k], orange[1]+directs[k+1]
				if row < 0 || row >= m || col < 0 || col >= n {continue}
				if grid[row][col] != 1 {continue}
				grid[row][col] = 2
				queue = append(queue, []int{row, col})
			}
		}
		if len(queue) > 0 {step++}
	}

	for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 { return -1 }
        }
    }
	return step
}
