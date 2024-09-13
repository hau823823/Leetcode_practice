package combinationSum

func CombinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	// tmp 用来临时存储当前的组合
	// nums 是候选数字数组即 candidates
	// idx 表示当前考虑的 nums 中的索引
	// tgt 是当前递归中需要达到的目标和
	var backtrack func([]int, []int, int, int)
	backtrack = func(tmp, nums []int, idx, tgt int) {
		if tgt < 0 || idx > len(nums) {
			return
		}

		if tgt == 0 {
			cpyTmp := make([]int, len(tmp))
			copy(cpyTmp, tmp)
			result = append(result, cpyTmp)
		}

		for i := idx; i < len(nums); i++ {
			tmp = append(tmp, nums[i])
			backtrack(tmp, nums, i, tgt - nums[i])
			tmp = tmp[:len(tmp) - 1]
		}
	}

	backtrack(make([]int, 0), candidates, 0, target)
    return result
}