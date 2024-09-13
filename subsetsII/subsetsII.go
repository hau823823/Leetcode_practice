package subsetsII

import (
	"sort"
)

func SubsetsWithDuSort(nums []int) [][]int {
	result := [][]int{{}}

	var backtrack func([]int, []int)
	backtrack = func(nums, cur []int) {
		for i := 0; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}

			cur = append(cur, nums[i])
			copy := append([]int{}, cur...)
			result = append(result, copy)

			backtrack(nums[i+1:], cur)
			cur = cur[:len(cur)-1]
		}
	}

	sort.Ints(nums)
	backtrack(nums, []int{})
	return result
}

func SubsetsWithDup(nums []int) [][]int {
	result := [][]int{}

	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	var backtrack func([]int)
	backtrack = func(cur []int) {
		cpy := append([]int{}, cur...)
		result = append(result, cpy)

		for k := range counts {
			// 维持子集的元素顺序以保证每个子集是排序的
			// 这通过确保不会向子集中添加小于当前子集中最后一个元素的新元素来实现
			// 避免了重复生成相同子集的情况
			if counts[k] == 0 || (len(cur) > 0 && cur[len(cur)-1] > k) {
				continue
			}
			counts[k]--
			cur = append(cur, k)
			backtrack(cur)
			counts[k]++
			cur = cur[:len(cur)-1]
		}
	}

	backtrack([]int{})
	return result
}
