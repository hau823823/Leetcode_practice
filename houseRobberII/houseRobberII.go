package houseRobberII

func Rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	} else if l == 2 {
		return max(nums[0], nums[1])
	}

	return max(robWithoutCircle(nums[1:]), robWithoutCircle(nums[:l - 1]))
}

func robWithoutCircle(nums []int) int {
	dp := make([]int, len(nums))
	m := max(nums[0], nums[1])
	dp[0], dp[1] = nums[0], m

	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i] + dp[i - 2], dp[i - 1])
		m = max(m, dp[i])
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}