package reverseInt

import (
	"math"
)

func Reverse(x int) int {
	set := make([]int, 0)
	for x != 0 {
		set = append(set, x%10)
		x /= 10
	}

	res := 0
	for i := len(set) - 1; i >= 0; i-- {
		res += set[i] * int(math.Pow10(len(set)-i-1))
	}

	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	return res
}

func ReverseOp(x int) int {
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x /= 10
	}
	
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	return res
}