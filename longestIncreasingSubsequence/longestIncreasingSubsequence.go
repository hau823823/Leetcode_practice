package longestIncreasingSubsequence

import (
	"fmt"
	"sort"
)

func LengthOfLIS1(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	res := 1
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			//fmt.Println("nums[",i,"]:", nums[i])
			//fmt.Println("nums[",j,"]:", nums[j])
			//fmt.Println(dp)
		}
		res = max(res, dp[i])
	}

	return res
}

func LengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums))
	res := 0

	for i := range nums {
		curMax := 0
		for j := 0; j < i+1; j++ {
			if nums[j] < nums[i] {
				curMax = max(curMax, dp[j])
			}
		}

		dp[i] = curMax + 1
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LengthOfLISSearch(nums []int) int {
	dp := make([]int, len(nums))
	counter := 0
	
	for _, num := range nums {
		pos := sort.SearchInts(dp[0:counter], num)
		dp[pos] = num
		if pos == counter {
			counter++
		}
		fmt.Println("pos: ", pos)
		fmt.Println("counter: ", counter)
		fmt.Println("dp: ", dp)
	}

	return counter
}