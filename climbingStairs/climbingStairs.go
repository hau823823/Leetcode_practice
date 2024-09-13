package climbingStairs

func ClimbStairs(n int) int {

	if n == 0 || n == 1 || n == 2 {
		return n
	}

	var steps [46]int
	steps[0] = 0
	steps[1] = 1
	steps[2] = 2

	for i := 3; i <= n; i++ {
		steps[i] = steps[i-1] + steps[i-2]
	}
	
	return steps[n]
}