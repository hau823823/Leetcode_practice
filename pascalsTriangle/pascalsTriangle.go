package pascalsTriangle

func Generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = append(res[0], 1)

	for i := 1; i < numRows; i++ {
		res[i] = append(res[i], res[i-1]...)
		res[i] = append(res[i], 0)

		for j := 1; j < len(res[i]); j++ {
			res[i][j] += res[i-1][j-1]
		}
	}

	return res
}

func GenerateOp(numRows int) [][]int {
	res := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1

		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}