package string2Integer

import (
	"math"
)

func MyAtoi(s string) int {

	if len(s) == 0 {
		return 0
	}

	b := []byte(s)
	b = append(b, 'n')

	plusOrMinus := 1
	nums := make([]int, 0)

	for i := 0; i < len(b); i++ {

		// 判斷是否是 空字元
		if b[i] == ' ' {
			continue
		}

		// 判斷是否有 負數
		if (i != len(s) - 1) && (b[i] == '-' || b[i] == '+') && (isDigital(b[i+1])) {
			if b[i] == '-' {
				plusOrMinus = plusOrMinus * -1
			}
			continue
		}

		// 判斷是否是 digital
		if !isDigital(b[i]) {
			break
		} else if isDigital(b[i]) {
			nums = append(nums, str2Digital(b[i]))

			if !isDigital(b[i+1]) {
				break
			}
		}
	}

	result := float64(0)
	for i := 0; i < len(nums); i++{
		digit := math.Pow(10, float64(len(nums) - 1 - i))

		result += float64(nums[i]) * digit
	}
	result = result * float64(plusOrMinus)

	if result > math.MaxInt32 {
		return math.MaxInt32
	} else if result < math.MinInt32 {
		return math.MinInt32
	}

	return int(result)
}

func isDigital(s byte) bool {
	if s < '0'|| s > '9' {
		return false
	}
	return true
}


func str2Digital(s byte) int {
	return int(s - '0')
}