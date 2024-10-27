package targetSum

// brute force
func FindTargetSumWays(nums []int, target int) int {
	count := 0

	var backtrack func(int, int)
	backtrack = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		backtrack(index+1, sum-nums[index])
		backtrack(index+1, sum+nums[index])
	}

	backtrack(0, 0)
	return count
}

// memorized