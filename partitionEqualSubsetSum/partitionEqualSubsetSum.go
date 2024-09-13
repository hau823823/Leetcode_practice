package partitionEqualSubsetSum

// time complexity: O(n ∗ sum(nums))
func CanPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum % 2 != 0 {
		return false
	}

	dp := make(map[int]bool)
	dp[0] = true

	for i := len(nums) - 1; i >= 0; i-- {
		nextDP := make(map[int]bool)
		for t := range dp {
			if t + nums[i] == sum / 2 {
				return true
			}
			nextDP[t + nums[i]] = true 
			nextDP[t] = true
		}
		dp = nextDP
	}
	
	return false
}

func CanPartitionOp(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum % 2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target + 1)
	dp[0] = true

	for _, num := range nums {
		for j := target; j >= num; j-- {
			if dp[j - num] {
				dp[j] = true
			}
		}
	}

	return dp[target]
}