package subset

import "fmt"

// Iterative
func Subsets(nums []int) [][]int {
	subsets := [][]int{{}}

	for _, num := range nums {
		fmt.Println(subsets)
		for _, subset := range subsets {
			fmt.Println("start: ", subset)
			sub := append(subset, num)
			fmt.Println("sub: ", sub)
			subset := append([]int{}, sub...)
			// 这一行代码创建了 sub 的一个副本
			// 确保将一个新的独立子集添加到子集列表中
			// 这一步是必需的，因为 Go 的切片是引用类型
			// 直接添加 sub 可能会导致未来修改影响到已添加的子集
			fmt.Println("subset: ", subset)
			subsets = append(subsets, subset)
			fmt.Println("subsets: ",subsets)
			fmt.Printf("\n")
		}
		fmt.Printf("---------------------------\n")
	}
	return subsets
}

// backtrack
func SubsetsBackTrack(nums []int) [][]int {
	result := make([][]int, 0)

	var backtrack func([]int, int)
	backtrack = func(current []int, start int) {
		subsetCopy := make([]int, len(current))
    	copy(subsetCopy, current)
		result = append(result, subsetCopy)

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(current, i+1)
			current = current[:len(current)-1]
		}
	}

	backtrack([]int{}, 0)
	return result
}