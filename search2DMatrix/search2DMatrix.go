package search2DMatrix

// original
// time complexity: O(m + log(n))
func SearchMatrix(matrix [][]int, target int) bool {
	tmp := []int{}
	for _, m := range matrix {
		tmp = append(tmp, m...)
	}

	left, right := 0, len(tmp) - 1
	for left <= right {
		mid := (left + right) / 2
		if tmp[mid] == target {
			return true
		} else if tmp[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

// optimize 
// time complexity: O(log(m * n))
func SearchMatrixOp(matrix [][]int, target int) bool {
	left, right := 0, len(matrix[0]) - 1
	up, down := 0, len(matrix) - 1

	for up <= down {
		rowMid := (up + down) / 2
		if matrix[rowMid][0] <= target && target <= matrix[rowMid][len(matrix[0]) - 1] {
			for left <= right {
				colMid := (left + right) / 2
				if matrix[rowMid][colMid] == target {
					return true
				} else if target < matrix[rowMid][colMid] {
					right = colMid - 1
				} else {
					left = colMid + 1
				}
			}
			break
		} else if target < matrix[rowMid][0] {
			down = rowMid - 1
		} else {
			up = rowMid + 1
		}
	}
	return false
}

// optimize
func SearchMatrixOp2(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	left, right := 0, rows * cols - 1

	for left <= right {
		mid := (left + right) / 2
		row := mid / cols
		col := mid % cols
		if target == matrix[row][col] {
			return true
		} else if target > matrix[row][col] {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}

	return false
}