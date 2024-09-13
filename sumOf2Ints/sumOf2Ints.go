package sumOf2Ints

func GetSum(a int, b int) int {
	sum := a ^ b
	co := a & b
	for co != 0 {
		co = co << 1
		sum,co = sum ^ co, sum & co
	}

	return sum
}

