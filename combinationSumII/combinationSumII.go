package combinationSumII

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)

	var backtrack func([]int, int, int)
	backtrack = func(cur []int, curSum, idx int) {
		if curSum > target {
			return
		}

		if curSum == target {
			cpy := append([]int{}, cur...)
			res = append(res, cpy)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if i > idx && candidates[i] == candidates[i-1]{
				continue
			}
			cur = append(cur, candidates[i])
			backtrack(cur, curSum + candidates[i], i + 1)
			cur = cur[:len(cur) - 1]
		}
	}

	backtrack([]int{}, 0, 0)
	return res
}
