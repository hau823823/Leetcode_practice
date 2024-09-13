package plusOne

import (
	"math"
)


func PlusOne(digits []int) []int {
    result := make([]int, 0)

	l := len(digits)
	var tmp uint64
	tmp = 0
	for i, v := range digits {
		tmp += uint64(v) * uint64(math.Pow10(l-i-1))
	}

	tmp += 1
	for tmp > 0 {
		v := int(tmp%10)
		tmp /= 10

		result = append([]int{v}, result...)
	}

	return result
}

func PlusOne2(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0

	}
	digits = append([]int{1}, digits...)
	return digits
}