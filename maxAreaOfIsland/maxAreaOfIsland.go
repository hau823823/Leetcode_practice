package maxAreaOfIsland

// dfs
func MaxAreaOfIsland(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	max := 0

	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return 0
		}

		grid[i][j] = 0
		return dfs(i+1, j) + dfs(i-1, j) + dfs(i, j+1) + dfs(i, j-1) + 1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			max = maxInt(max, dfs(i, j))
		}
	}

	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

