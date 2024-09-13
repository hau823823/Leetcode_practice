package rotateImage


func Rotate(matrix [][]int) {
	l := len(matrix)

	for i := 0; i < l/2; i++ {
		for j := i; j < l-i-1; j++ {
			matrix[i][j], matrix[j][l-i-1] = matrix[j][l-i-1], matrix[i][j]
			matrix[i][j], matrix[l-i-1][l-j-1] = matrix[l-i-1][l-j-1], matrix[i][j]
			matrix[i][j], matrix[l-j-1][i] = matrix[l-j-1][i], matrix[i][j]
		}
	}
}

func Rotate2(matrix [][]int) {
	// step1:
	// 將 column 豎直改成 row
	// 沿著對角線 swap
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// step2:
	// 將 row 進行反轉
	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[i][l-j-1] = matrix[i][l-j-1], matrix[i][j]
		}
	}
}
