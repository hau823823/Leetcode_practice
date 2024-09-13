package powerOfThree

// 3 的次方
func IsPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n%3 != 0 || n <= 0 {
		return false
	}

	return IsPowerOfThree(n/3)
}