package search2dMatrixII

// wrong answer
// 原本設想可以分區查找
// 代码试图通过将矩阵划分为四个区域来缩小搜索范围。
// 它选择矩阵的四个角点（min, sec, thr, max），然后根据 target 与这些点的比较来调整搜索范围。
// 然而，这种方法的逻辑比较复杂，而且可能无法正确处理所有情况。
// 例如，当 target 小于 matrix[min[0]][min[1]] 时，代码减小了 maxCol 和 maxRow，
// 但实际上 target 仍可能存在于当前行的右边或当前列的下方
func SearchMatrixWrong(matrix [][]int, target int) bool {
	minCol, maxCol := 0, len(matrix)-1
	minRow, maxRow := 0, len(matrix[0])-1

	for minRow <= maxRow && minCol <= maxCol {
		min := []int{(minCol + maxCol) / 2, (minRow + maxRow) / 2}
		sec := []int{(minCol + maxCol) / 2, maxRow}
		thr := []int{maxCol, (minRow + maxRow) / 2}
		max := []int{maxCol, maxRow}
		if target == matrix[min[0]][min[1]] || target == matrix[sec[0]][sec[1]] || target == matrix[thr[0]][thr[1]] || target == matrix[max[0]][max[1]] {
			return true
		} else if target < matrix[min[0]][min[1]] {
			maxCol, maxRow = min[0], min[1]
		} else if target < matrix[sec[0]][sec[1]] {
			maxCol, maxRow = sec[0], sec[1]
			minRow = thr[1]
		} else if target < matrix[thr[0]][thr[1]] {
			maxCol, maxRow = thr[0], thr[1]
			minCol = sec[0]
		} else {
			minCol, minRow = min[0], min[1]
			maxCol, maxRow = max[0], max[1]
		}
		//fmt.Println(minCol, maxCol)
		//fmt.Println(minRow, maxRow)
	}

	return false
}


func SearchMatrix(matrix [][]int, target int) bool {
	maxRow, maxCol := len(matrix), len(matrix[0])
	row, col := 0, maxCol - 1

	for row < maxRow && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if target < matrix[row][col] {
			col--
		} else {
			row++
		}
	}

	return false
}