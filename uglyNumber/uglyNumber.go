package uglyNumber

func IsUgly(n int) bool {

	if n < 0 {
		return false
	}

	divide := func(n int, divider int) int {
		for n % divider == 0 {
			n /= divider
		}
		return n
	}

	n = divide(n, 2)
	n = divide(n, 3)
	n = divide(n, 5)

	return n == 1
}
