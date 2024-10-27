package targetSum

import "fmt"

// brute force
func FindTargetSumWaysBF(nums []int, target int) int {
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
func FindTargetSumWaysMemo(nums []int, target int) int {
	var backtrack func(int, int, map[string]int) int
	backtrack = func(idx, target int, memo map[string]int) int {
		key := fmt.Sprint(idx) + "*" + fmt.Sprint(target)

		if val, exists := memo[key]; exists {
			return val
		}

		if idx == len(nums) {
			if target == 0 {
				return 1
			}
			return 0
		}

		res := backtrack(idx + 1, target + nums[idx], memo) + backtrack(idx + 1, target - nums[idx], memo)
		memo[key] = res
		return res
	}

	return backtrack(0, target, make(map[string]int))
}
