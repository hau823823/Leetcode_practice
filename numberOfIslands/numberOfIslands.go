package numberOfIslands

// bfs
// 通过从一个岛屿部分开始，然后逐渐向外扩展到所有相连的部分
// 并且在访问过程中将它们标记为 '0'，算法能够确保每个岛屿只被计数一次。
// 按层次扩展岛屿的边界
func NumIslandsBFS(grid [][]byte) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m,n := len(grid), len(grid[0])
	islands := 0
	offsets := []int{0, 1, 0, -1, 0} // 右、上、左、下 (兩兩一組：0,1 、 1,0 、 0,-1 、 -1,0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1'{
				islands++
				grid[i][j] = '0'

				queue := [][]int{{i, j}}
				for len(queue) > 0 {
					p := queue[0]
					queue = queue[1:]
					for k := 0; k < 4; k++{
						row, col := p[0] + offsets[k], p[1] + offsets[k+1]

						// 檢查是否在陣列內
						if row < 0 || row >= m || col < 0 || col >= n {
							continue
						}
						// 檢查是否是島嶼
						if grid[row][col] == '1' {
							grid[row][col] = '0'
							queue = append(queue, []int{row, col})
						}
					}
				}
			}
		}
	}

	return islands
}

// dfs
// 深入探索每个岛屿的每个部分，直到没有更多相连的陆地
func NumIslandsDFS(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	islands := 0

	var eraseIslands func(int, int)
	eraseIslands = func(i, j int){
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'
		eraseIslands(i-1, j)
        eraseIslands(i+1, j)
        eraseIslands(i, j-1)
        eraseIslands(i, j+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				islands++
				eraseIslands(i, j)
			}
		}
	}

	return islands
}