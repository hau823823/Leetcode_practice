package setMatrixZero

func SetZeroes(matrix [][]int)  {
    rmap := make(map[int]int8) 
	cmap := make(map[int]int8)

	for r, row := range matrix {
		for c, v := range row {
			if v == 0 {
				rmap[r] = 0
				cmap[c] = 0
			}
		}
	}

	for r, row := range matrix {
		_, ok := rmap[r]
		if ok {
			for i, _ := range row {
				row[i] = 0
			}
		} else {
			for i, _ := range row {
				_, ok := cmap[i]
				if ok {
					row[i] = 0
				}
			}
		}
	}
}

func SetZeroesOp(matrix [][]int) {
	firstRowZero := false
	firstColZero := false

	// 使用第一行和第一列來標記其餘行列是否需要置零
	for r, row := range matrix {
		for c, v := range row {
			if v == 0 {
				if r == 0 {
					firstRowZero = true
				}

				if c == 0 {
					firstColZero = true
				}

				matrix[r][0] = 0
				matrix[0][c] = 0
			}
		}
	}

	// 置零除第一行以外的行列
	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if firstRowZero {
		for r, _ := range matrix[0]{
			matrix[0][r] = 0
		}
	}

	if firstColZero {
		for c, _ := range matrix {
			matrix[c][0] = 0
		}
	}

}