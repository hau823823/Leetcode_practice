package factorialTrailingZeroes

import "fmt"

func TrailingZeroes(n int) int {
	result := n
	n -= 1
	count := 0
	for n > 0 {
		result *= n
		n--
		if result%10 == 0 {
			result /= 10
			count++
		}
		fmt.Println(result)
	}

	return count
}

func TrailingZeroesOp(n int) int {
	ans := 0
	tmp := 5

	for n/tmp != 0 {
		ans += n/tmp
		tmp *= 5
	}

	return ans
}
