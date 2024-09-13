package uniquePath

import "fmt"

func UniquePaths(m int, n int) int {
	stepSet := make([][]int, m)

	// left
	for i := 0; i < n; i++ {
		stepSet[0] = append(stepSet[0], 1)
	}

	// down
	for i := 1; i < m; i++ {
		stepSet[i] = append(stepSet[i], 1)
	}

	// count
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			stepSet[i] = append(stepSet[i], stepSet[i-1][j] + stepSet[i][j-1])
		}
	}
	fmt.Println(stepSet)
	return stepSet[m-1][n-1]
}
