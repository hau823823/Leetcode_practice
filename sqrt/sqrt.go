package sqrt

import "math"
//import "fmt"


// use pack
func MySqrt(x int) int {
    return int(math.Sqrt(float64(x)))
}

// brute
func MySqrtBrute(x int) int {
	res := 0
	sqrt := 0
	for sqrt <= x {
		res += 1
		sqrt = res * res
		if sqrt > x {
			return res-1
		}
	}
	return res
}

// binary search
func MySqrtBS(x int) int {
	left, right, mid := 1, x, 0
    for left <= right {
        mid = left + (right - left)/2
        sqrt := x / mid
        if mid == sqrt {
            return mid
        } else if mid > sqrt {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return right
}