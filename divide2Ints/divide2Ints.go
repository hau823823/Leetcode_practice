package divide2Ints

import "math"

// origin
func Divide(dividend int, divisor int) int {
	return dividend / divisor
}

func Divide2(dividend int, divisor int) int {
	// single special case that would cause overflow
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	quotient := 0
	divid := abs(dividend)
	divis := abs(divisor)
	
	for divid >= divis {
		sub := divis
		add := 1
		for divid >= sub << 1 {
			sub <<= 1
			add <<= 1
		}
		divid -= sub
		quotient += add
	}

	if (dividend < 0) != (divisor < 0) {
		return -quotient
	}

	return quotient
}


func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}